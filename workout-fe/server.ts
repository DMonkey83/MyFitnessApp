import { createServer } from 'http';
import { parse } from 'url';
import { createProxyMiddleware } from 'http-proxy-middleware';
import { NextApiHandler } from 'next';

import next from 'next';

const dev = process.env.NODE_ENV !== 'production';
const app = next({ dev });
const handle: NextApiHandler = app.getRequestHandler();

app.prepare().then(() => {
  createServer((req, res) => {
    const parsedUrl = parse(req.url!, true);
    const { pathname } = parsedUrl;

    // Define your API proxy target
    const apiProxy = createProxyMiddleware('/api', {
      target: 'http://localhost:8080', // Replace with your backend URL
      changeOrigin: true,
    });

    if (pathname && pathname.startsWith('/api')) {
      // Use the API proxy for requests matching the '/api' path
      apiProxy(req, res);
    } else {
      // Handle other requests as usual
      handle(req, res, parsedUrl);
    }
  }).listen(3000, (err: any) => {
    if (err) throw err;
    console.log('> Ready on http://localhost:3000');
  });
});
