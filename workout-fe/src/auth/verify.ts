import { V2 } from 'paseto'

export const verify = (token: string, publicKey: string) => {
  V2.verify(
    token, publicKey)
}
