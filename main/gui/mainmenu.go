components {
  id: "gui"
  component: "/main/gui/mainmenu.gui"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sound"
  type: "sound"
  data: "sound: \"/main/sounds/click.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 0.1\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "soundmusic"
  type: "sound"
  data: "sound: \"/main/sounds/menumusic.ogg\"\n"
  "looping: 1\n"
  "group: \"master\"\n"
  "gain: 0.031622775\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
