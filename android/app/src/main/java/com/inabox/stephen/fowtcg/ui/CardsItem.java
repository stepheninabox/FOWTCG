package com.inabox.stephen.fowtcg.ui;

import android.text.Html;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;
import butterknife.ButterKnife;
import butterknife.InjectView;
import com.codesmyth.droidcook.common.widget.ViewBinder;
import com.inabox.stephen.fowtcg.R;

public class CardsItem extends ViewBinder<CardsAdapter> implements View.OnClickListener {

  @InjectView(R.id.name)
  TextView mName;

  @InjectView(R.id.description)
  TextView mDescription;

  public CardsItem(ViewGroup parent) {
    super(LayoutInflater.from(parent.getContext()).inflate(R.layout.cards_item, parent, false));
    ButterKnife.inject(this, itemView);
    itemView.setOnClickListener(this);
  }

  @Override
  public void bind(CardsAdapter adapter, int pos) {
    mName.setText(adapter.getCursor().getString("name"));
    mDescription.setText(Html.fromHtml(adapter.getCursor().getString("description")));
  }

  @Override
  public void onClick(View v) {

  }
}
