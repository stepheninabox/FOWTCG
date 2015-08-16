package com.inabox.stephen.fowtcg;

import android.database.sqlite.SQLiteOpenHelper;
import android.net.Uri;
import android.support.annotation.NonNull;
import com.codesmyth.droidcook.common.content.Content;

public class AppContent extends Content {

  public static final String AUTHORITY = "com.inabox.stephen.fowtcg";
  public static final Uri CONTENT_URI = Uri.parse("content://" + AUTHORITY);
  public static final Uri CARDS_URI = CONTENT_URI.buildUpon().appendPath("cards").build();

  @NonNull
  @Override
  public SQLiteOpenHelper getDatabase() {
    return new AppDatabase(getContext());
  }

  @NonNull
  @Override
  public String getTableName(Uri uri) {
    return uri.getLastPathSegment();
  }
}
