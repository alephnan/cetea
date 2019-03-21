document.addEventListener("DOMContentLoaded", () => {
  window.vanillaTip.init();
	apiHealthCheck();

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
      fetch("http://localhost:8080/api/authorization", {
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


// https://github.com/nahojd/vanilla-tooltip
(function() {
	function Tip() {
		function isTip(element) {
			return element.classList.contains('tip');
		}

		function isInPopover(element) {
			if (!element || element.tagName.toLowerCase() === 'body')
				return false;

			if (element.classList.contains('popover-clone'))
				return true;

			return isInPopover(element.parentElement);
		}

		function hidePopovers() {
			const popovers = document.querySelectorAll('.popover-clone');
			for (let i = 0; i < popovers.length; i++) {
				document.body.removeChild(popovers[i]);
			}
			const links = document.querySelectorAll('.tip');
			for (let i = 0; i< links.length; i++) {
				links[i].dataOpenPopover = null;
			}
		}

		function toggleTooltip(target) {
			//Remove open popover if exists
			if (target.dataOpenPopover) {
				document.body.removeChild(target.dataOpenPopover);
				target.dataOpenPopover = null;
				return false;
			}

			//Or create and show popover
			const originalPopover = target.nextElementSibling;
			if (!originalPopover)
				return false;

			const popover = originalPopover.cloneNode(true);
			popover.classList.add('popover-clone');
			document.body.appendChild(popover);
			setPosition(popover, target);
			target.dataOpenPopover = popover;

			return false;
		}

		function setPosition(popover, target) {
      const margin = 10;
			const scrollY = window.pageYOffset || document.documentElement.scrollTop || 0;
			const scrollX = window.pageXOffset || document.documentElement.scrollLeft || 0;
			const targetRect = target.getBoundingClientRect();
			const targetWidth = targetRect.width || targetRect.left - targetRect.right;
			const targetHeight = targetRect.height || targetRect.top - targetRect.bottom;
			//Position the popup window
			popover.style.display = 'block';
			popover.style.position = 'absolute';
			let left = targetWidth / 2 + targetRect.left + scrollX - popover.clientWidth / 2;
			if (left < 10) { left = 10; }
			popover.style.left = left + 'px';
			popover.style.bottom = document.documentElement.clientHeight - targetRect.top - scrollY + targetHeight/2 + margin + 'px';
			popover.style.top = 'inherit';
			popover.style.right = 'inherit';
			popover.style.zIndex = 10000;

			//Position the arrow over the clicked element
			if (popover.querySelector) {
				const arrow = popover.querySelector('.arrow');
				arrow.style.left = targetRect.left - left + arrow.offsetWidth / 2 + 'px';
			}
		}

		function init() {
			document.body.addEventListener('click', function (e) {
				const target = e.target || e.srcElement;
				if (isTip(target) || isInPopover(target))
					return;

				hidePopovers();
			});
      window.addEventListener('resize', hidePopovers);
    }

		return {
      click: toggleTooltip,
      hide: hidePopovers,
			init: init
		};
	};

  window.vanillaTip = Tip();
})();

/**
 * Pings API to check that it's up.
 */
function apiHealthCheck() {
	fetch("http://localhost:8080/api/health",  {
		method: "GET",
		cache: "no-cache",
		credentials: "same-origin",
	}).then(response => {
		// TODO: alert in the UI
		console.log(response.json())
	});
}