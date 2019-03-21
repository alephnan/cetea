<template>
  <div id="main">
    <div id="sidenav" class="hidden">
      <div id="sidenav-header">
        <div id="sidenav-header-title">Projects</div>
        <div id="sidenav-header-button">
          <div class="button-container">
            <button href="#" class="button" id="new-project">
              New
            </button>
          </div>
        </div>
      </div>
      <div>
        <div id="sidenav-projectlist-spinner-container">
          <div id="sidenav-projectlist-spinner" class="lds-roller"><div></div><div></div><div></div><div></div><div></div><div></div><div></div><div></div></div>
        </div>
        <div id="sidenav-projectlist-container" class="hidden">
          <ul id="sidenav-projectlist">
            <!-- TODO: handle too long name. truncate, use ellipsis. -->
            <li>
              <span class="tip" onmouseover="return vanillaTip.click(this);"  onmouseout="return vanillaTip.hide()">tooltip</span>
              <div class="popover top">
                <div class="arrow"></div>
                <div class="popover-content">
                  organization: Google
                  project_name: abc
                  project_id: gfjfj
                  project_number: jgkg
                  created on: jghgfg
                </div>
              </div>
            </li>
            <li>
              <span class="tip" onmouseover="return vanillaTip.click(this);"  onmouseout="return vanillaTip.hide()">Theon56k</span>
              <div class="popover top">
                <div class="arrow"></div>
                <div class="popover-content">
                  ABCfkjfgkfgjkgjgf
                </div>
              </div>
            </li>
            <li><a href="">goolf</a></li>
            <li><a href="">appengine</a></li>
            <li><a href="">hyperapp</a></li>
            <li><a href="">himalayan</a></li>
            <li><a href="">gradient</a></li>
            <li><a href="">zamboni</a></li>
            <li><a href="">parthenon</a></li>
            <li><a href="">atlas</a></li>
          </ul>
          <a href="" id="show-more">
            Show more
          </a>
        </div>
      </div>
    </div>
    <div id="content">
      <div id="orgline">organization: <span id="organization">google.com</span></div>
      <div id="project">
        <span id="projectlabel">project:</span>
        <span id="projectname">my-first-project</span>
        <span id="projectid">( uplifted-scout-234505 )</span>
      </div>
    </div>
  </div>
</template>