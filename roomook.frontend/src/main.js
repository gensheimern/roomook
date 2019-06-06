import Vue from 'vue';
import App from './App.vue'

import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import VueCookies from 'vue-cookies'
import Login from './components/Login'
import VueRouter from 'vue-router'
import CalendarWrapper from './components/CalendarWrapper'
import VueNativeSock from 'vue-native-websocket'
import auth from './middleware/auth'
import Tablet from './components/TabletUI'



Vue.use(Vuetify, {
  theme: {
    primary: '#274d6f',
    secondary: '#b0bec5',
    accent: '#8c9eff',
    error: '#b71c1c'
  },
  options: {
    customProperties: true
  }
})
Vue.use(VueNativeSock, 'wss://roomook-be.jacob.run:8443/ws/', { format: 'json' })
Vue.use(VueCookies);
Vue.use(require('vue-moment'));

Vue.use(VueRouter);

Vue.config.productionTip = false

$cookies.config('365d')

Vue.prototype.$url = (process.env.VUE_APP_AUTH_AS_MIDDLEWARE != 0) ? process.env.VUE_APP_ROOMOOK_API_WITH_AUTH : process.env.VUE_APP_ROOMOOK_API


const router = new VueRouter({
  routes: [
    {
      path: '/',
      name: 'CalenderView',
      component: CalendarWrapper,
      meta: {
        middleware: auth,
      },
      mode: 'history',


    },

    {
      path: "/tablet/:id",
      name: "tablet",
      component: Tablet
    },

    {
      path: "/login",
      name: "login",
      component: Login
    }
  ]
})


function nextFactory(context, middleware, index) {
  const subsequentMiddleware = middleware[index];
  // If no subsequent Middleware exists,
  // the default `next()` callback is returned.
  if (!subsequentMiddleware) return context.next;

  return (...parameters) => {
    // Run the default Vue Router `next()` callback first.
    context.next(...parameters);
    // Then run the subsequent Middleware with a new
    // `nextMiddleware()` callback.
    const nextMiddleware = nextFactory(context, middleware, index + 1);
    subsequentMiddleware({ ...context, next: nextMiddleware });
  };
}
router.beforeEach((to, from, next) => {
  if (to.meta.middleware) {
    const middleware = Array.isArray(to.meta.middleware)
      ? to.meta.middleware
      : [to.meta.middleware];

    const context = {
      from,
      next,
      router,
      to,
    };
    const nextMiddleware = nextFactory(context, middleware, 1);

    return middleware[0]({ ...context, next: nextMiddleware });
  }

  return next();
});

new Vue({

  router: router,
  render: (h) => h(App),
}).$mount('#app')
