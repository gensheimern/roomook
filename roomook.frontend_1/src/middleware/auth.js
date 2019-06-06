export default function auth({ next, router }) {
    if (!localStorage.getItem('Authentication')) {
      return router.push({ name: 'login' });
    }
  
    return next();
  }