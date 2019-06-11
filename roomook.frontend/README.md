# Roomook Frontend

## Getting started

Install dependencies

``` 
npm install
```
## Lints and fixes files

```
npm run lint
```

## Tests

Tests are not implemented yet. To run tests run:

```
npm run test
```

## Compiles and hot-reloads for development
```
npm run serve
```

## Compiles and minifies for production
```
npm run build
```

## HTTP Requests

Since we are using our Authentication-Proxy (See [more](https://igitlab.jacob.de/development/frontend/authentication-proxy)), HTTP Requests needs to be send to a different url.

For developing you can enable/disable the authentication in the ***.env.development*** file, locatet in the root folder.


```
VUE_APP_AUTH_AS_MIDDLEWARE=1
VUE_APP_ROOMOOK_API=http://localhost:3000
VUE_APP_ROOMOOK_API_WITH_AUTH=""
```

### No auth

To disable authentication set *VUE_APP_AUTH_AS_MIDDLEWARE* to 0
Then the HTTP Request URL will be *VUE_APP_ROOMOOK_API*.

HTTP Requests will go directly to your Backend now.


### With auth

To disable authentication set *VUE_APP_AUTH_AS_MIDDLEWARE* to 1
Then the HTTP Request URL will be *VUE_APP_ROOMOOK_API_WITH_AUTH*. Leave it blank for localhost.

HTTP Requests will go to the authentication proxy and he will pass it to the backend.


### Development mode

When you run the frontend in dev mode (npm run serve), webpack is used to deliver the website. 

The file ***vue.config.js***, which is locatet in the root foleder, handles the redirects to the authentication proxy. It also solves some CORS problems with vue.


```
module.exports = {
  devServer: {
    proxy: {
    
      ... 
      
      '^/booking': {
        target: 'https://auth.jacob.run/proxy/roomookBackend/q/',
        ws: true,
      },
      
      ...

    }
  }
}
```

### Production mode

When you run the frontend in prod mode or you start the docker container, NGINX is used to deliver the website.

The file ***default.conf***, which is located in *nginx/default.conf*, handles the redirects to the authentication proxy.


```
...

  location /booking {
    proxy_pass https://auth.jacob.run/proxy/roomookBackend/q/booking;
  }

... 
```