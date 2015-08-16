package com.inabox.stephen.fowtcg.ui;

import android.content.Context;
import android.database.Cursor;
import android.os.Bundle;
import android.support.v4.content.Loader;
import android.view.ViewGroup;
import com.codesmyth.droidcook.common.content.Query;
import com.codesmyth.droidcook.common.widget.RecyclerCursorAdapter;
import com.codesmyth.droidcook.common.widget.ViewBinder;
import com.inabox.stephen.fowtcg.AppContent;
import com.inabox.stephen.fowtcg.CardsItem;

public class CardsAdapter extends RecyclerCursorAdapter<CardsAdapter> {

  public CardsAdapter(Context context) {
    super(context);
  }

  @Override
  public ViewBinder<CardsAdapter> onCreateViewHolder(ViewGroup parent, int viewType) {
    return new CardsItem(parent);
  }

  @Override
  public Loader<Cursor> onCreateLoader(int id, Bundle args) {
    return Query.builder()
        .select("_id", "name", "description")
        .from(AppContent.CARDS_URI)
        .cursorLoader(getContext());
  }
}
