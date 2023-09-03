package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/DMonkey83/MyFitnessApp/workout-be/db/mock"
	db "github.com/DMonkey83/MyFitnessApp/workout-be/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/workout-be/token"
	"github.com/DMonkey83/MyFitnessApp/workout-be/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetPlanAPI(t *testing.T) {
	plan := randomWorkoutPlan(t)
	testCases := []struct {
		name          string
		plan_id       int64
		username      string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:     "OK",
			plan_id:  plan.PlanID,
			username: plan.Username,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, plan.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.GetPlanParams{
					PlanID:   plan.PlanID,
					Username: plan.Username,
				}
				store.EXPECT().
					GetPlan(gomock.Any(), gomock.Eq(arg)).
					Times(1).Return(plan, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchWorkoutPlan(t, recorder.Body, plan)
			},
		},
		{
			name:     "Not Found",
			plan_id:  plan.PlanID,
			username: plan.Username,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, plan.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.GetPlanParams{
					PlanID:   plan.PlanID,
					Username: plan.Username,
				}
				store.EXPECT().
					GetPlan(gomock.Any(), gomock.Eq(arg)).
					Times(1).Return(db.Workoutplan{}, db.ErrRecordNotFound)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:     "Internal Error",
			plan_id:  plan.PlanID,
			username: plan.Username,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, plan.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.GetPlanParams{
					PlanID:   plan.PlanID,
					Username: plan.Username,
				}
				store.EXPECT().
					GetPlan(gomock.Any(), gomock.Eq(arg)).
					Times(1).Return(db.Workoutplan{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:     "Invalid ID",
			plan_id:  0,
			username: plan.Username,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, plan.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.GetPlanParams{
					PlanID:   plan.PlanID,
					Username: plan.Username,
				}
				store.EXPECT().
					GetPlan(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/workout-plan/%d/%s", tc.plan_id, tc.username)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreatePlanAPI(t *testing.T) {
	plan := randomWorkoutPlan(t)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"plan_name":   plan.PlanName,
				"username":    plan.Username,
				"difficulty":  plan.Difficulty,
				"goal":        plan.Goal,
				"description": plan.Description,
				"is_public":   plan.IsPublic,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, plan.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreatePlanParams{
					Username:    plan.Username,
					PlanName:    plan.PlanName,
					Description: plan.Description,
					Goal:        plan.Goal,
					IsPublic:    plan.IsPublic,
					Difficulty:  plan.Difficulty,
				}

				store.EXPECT().
					CreatePlan(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(plan, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchWorkoutPlan(t, recorder.Body, plan)
			},
		},
		{
			name: "NoAuthorization",
			body: gin.H{
				"plan_name":   plan.PlanName,
				"username":    plan.Username,
				"difficulty":  plan.Difficulty,
				"goal":        plan.Goal,
				"description": plan.Description,
				"is_public":   plan.IsPublic,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreatePlan(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"plan_name":   plan.PlanName,
				"username":    plan.Username,
				"difficulty":  plan.Difficulty,
				"goal":        plan.Goal,
				"description": plan.Description,
				"is_public":   plan.IsPublic,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, plan.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreatePlan(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Workoutplan{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidDifficulty",
			body: gin.H{
				"plan_name":   plan.PlanName,
				"username":    plan.Username,
				"difficulty":  "Invalid",
				"goal":        plan.Goal,
				"description": plan.Description,
				"is_public":   plan.IsPublic,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, plan.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreatePlan(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidGoal",
			body: gin.H{
				"plan_name":   plan.PlanName,
				"username":    plan.Username,
				"difficulty":  plan.Difficulty,
				"goal":        "invalid",
				"description": plan.Description,
				"is_public":   plan.IsPublic,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, plan.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreatePlan(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidVisibility",
			body: gin.H{
				"plan_name":   plan.PlanName,
				"username":    plan.Username,
				"difficulty":  plan.Difficulty,
				"goal":        plan.Goal,
				"description": plan.Description,
				"is_public":   "invalid",
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, plan.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreatePlan(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/api/workout-plan"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomWorkoutPlan(t *testing.T) (workoutPlan db.Workoutplan) {
	user, _ := randomUser(t)
	return db.Workoutplan{
		PlanID:      int64(util.GetRandomAmount(1, 1000)),
		Username:    user.Username,
		PlanName:    util.GetRandomUsername(8),
		Description: util.GetRandomUsername(20),
		Goal:        db.WorkoutgoalenumBuildMuscle,
		IsPublic:    db.VisibilityPrivate,
		Difficulty:  db.DifficultyAdvanced,
	}
}

func requireBodyMatchWorkoutPlan(t *testing.T, body *bytes.Buffer, plan db.Workoutplan) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotPlan db.Workoutplan
	err = json.Unmarshal(data, &gotPlan)
	require.NoError(t, err)
	require.Equal(t, plan, gotPlan)
}
