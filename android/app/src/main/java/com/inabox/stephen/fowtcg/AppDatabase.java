package com.inabox.stephen.fowtcg;

import android.content.Context;
import com.codesmyth.droidcook.common.database.Database;

public class AppDatabase extends Database {
  public AppDatabase(Context context) {
    super(context, "app.db", 1, R.raw.cards);
  }
}
