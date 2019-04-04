
import _Vue from "vue"; // <-- notice the changed import

export default function GoogleSignin(Vue: typeof _Vue, options?: any): void {
  let auth2Promise: any = null;
  Vue.prototype.$getGapi = () => {
    return new Promise((resolve) => {
      if (!auth2Promise) {
        auth2Promise = new Promise(resolve => {
          const script = document.createElement('script');
          script.setAttribute('src', "https://apis.google.com/js/platform.js");
          script.async = true;
          script.onload = () => {
            resolve();
          };
          document.getElementsByTagName('head')[0].appendChild(script);
        }).then(_ => {
          return new Promise(resolve => {
            (window as any).gapi.load('auth2', () => {
              const authOptions = {
                client_id: options.clientId
              };
              resolve((window as any).gapi.auth2.init(authOptions));
            });
          });
        });
      }
      resolve(auth2Promise);
    })
  }
}