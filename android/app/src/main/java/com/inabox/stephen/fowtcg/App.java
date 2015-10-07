package com.inabox.stephen.fowtcg;

import android.app.Application;
import go.game.Game;

public class App extends Application {
  @Override
  public void onCreate() {
    try {
      Game.OpenDB(getFilesDir().getAbsolutePath());
    } catch (Exception e) {
      e.printStackTrace();
    }
  }
}
