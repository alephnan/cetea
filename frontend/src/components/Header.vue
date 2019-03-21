
<template>
  <div id="header">
    <div id="top-bar">
      <div id="top-bar-spacer"></div>
      <div id="top-bar-right">
        <i class="material-icons topnav-button" title="Search">search</i>
        <i class="material-icons topnav-button" title="Dark mode">brightness_4</i>
        <i class="material-icons topnav-button" title="Products">apps</i>
        <a id="login-button">Sign in</a>
      </div>
    </div>
    <div id="second-bar">
      <ul>
        <li>All</li>
        <li>Images</li>
        <li>Woosh</li>
      </ul>
    </div>
  </div>
</template>

<style scoped lang="less">
#header {
  padding-left: 150px;
  padding-top: 20px;
  font-size: 14pt;
  border-bottom: 1px solid rgb(235, 235, 235);
}
#top-bar {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}
#top-bar-spacer {
  flex-grow: 99;
}
#top-bar-right {
  flex-grow: 1;
  padding-right: 20px;
  display: flex;
  align-items: center;
}

#second-bar {
  display: none;
} 
/*
#second-bar ul {
  padding-left: 0px;
  padding-top: 0px;
  margin: 0px;
  margin-top: 24px;
  padding-bottom: 0px;
}
#second-bar li {
  display:inline;
  padding-bottom: 12px;
  padding-right: 16px;
  padding-left: 16px;
  font-size: 13px;
  color: rgb(102, 102, 102);
  font-family: Roboto, Arial, sans-serif;
} */


#login-button {
  padding-top: 9px;
  padding-bottom: 9px;
  padding-left: 23px;
  padding-right: 23px;
  font-size: 14px;
  background-color: rgb(43, 125, 233);
  color: white;
  border-bottom-left-radius: 4px;
  border-bottom-right-radius: 4px;
  border-top-left-radius: 4px;
  border-top-right-radius: 4px;
  font-family: "Google Sans",Roboto,RobotoDraft,Helvetica,Arial,sans-serif;
  margin-left: 10px;
}
#login-button:hover {
  cursor: pointer;
  box-shadow: 1px 4px 5px 1px rgba(0,0,0,0.1);
}
.material-icons.topnav-button {
  color: rgb(95, 99, 104);
  padding: 8px;
}
.material-icons.topnav-button:hover {
  background-color: #ebecec;
  border-radius: 50%;
  cursor: pointer;
}
</style>