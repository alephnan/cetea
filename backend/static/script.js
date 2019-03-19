document.addEventListener("DOMContentLoaded", () => {
  const login = document.getElementById("login-button");
  login.onclick = () => {
    // https://developers.google.com/identity/sign-in/web/reference#gapiauth2offlineaccessoptions
    // const prompt = "select_account";
    const prompt = "consent;"
    auth2.grantOfflineAccess().then(response => {
      if(!response.code) {
        console.log("Error.")
        return;
      }
      document.getElementById('sidenav').classList.remove("hidden");

      // https://developers.google.com/identity/sign-in/web/reference#googleusergetid
      const googleUser = auth2.currentUser.get();

      console.log("Verifying with backend.");
      const {id_token} = googleUser.getAuthResponse();
      fetch("http://localhost:8080/authorization", {
          method: "POST",
          cache: "no-cache",
          credentials: "same-origin",
          headers: {
            "Content-Type": "application/json",
            "X-Requested-With": "XMLHttpRequest"
          },
          body: JSON.stringify({
            "code": response.code,
            // TODO: verify id_token on server
            "id_token": id_token
          }),
      }).then(response => {
        response.json().then(json => {
          const projectNames = json.projects;
          const newUl = document.createElement('ul');
          newUl.id = "sidenav-projectlist";
          for(let i = 0 ; i < projectNames.length; i++) {
            const item = document.createElement('li');
            const a = document.createElement('a');
            a.setAttribute("href", "")
            a.appendChild(document.createTextNode(projectNames[i]));
            item.appendChild(a);
            // https://coderwall.com/p/o9ws2g/why-you-should-always-append-dom-elements-using-documentfragments
            newUl.appendChild(item);
          }
          const frag = document.createDocumentFragment();
          frag.appendChild(newUl);
          const ul = document.getElementById("sidenav-projectlist");
          ul.parentNode.replaceChild(frag, ul);

          document.getElementById('sidenav-projectlist-spinner-container').classList.add("hidden");
          document.getElementById('sidenav-projectlist-container').classList.remove("hidden");
        });
        // TODO: update login "verifying" state.
      });

      const profile = googleUser.getBasicProfile();
      const email = profile.getEmail();
      // BasicProfile.getId()
      // BasicProfile.getName()
      // BasicProfile.getGivenName()
      // BasicProfile.getFamilyName()
      // BasicProfile.getImageUrl()
      // BasicProfile.getEmail();
      login.innerText = email + " ( verifying )";
    });
  };
});