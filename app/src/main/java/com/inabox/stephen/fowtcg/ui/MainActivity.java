package com.inabox.stephen.fowtcg.ui;

import android.app.Activity;
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.DefaultItemAnimator;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.Menu;
import android.view.MenuItem;
import butterknife.ButterKnife;
import butterknife.InjectView;
import com.inabox.stephen.fowtcg.R;


public class MainActivity extends AppCompatActivity {

  @InjectView(R.id.recycler)
  RecyclerView mRecycler;

  CardsAdapter mAdapter;

  @Override
  protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.activity_main);
    ButterKnife.inject(this);

    LinearLayoutManager lm = new LinearLayoutManager(this);
    mRecycler.setLayoutManager(lm);
    mRecycler.setItemAnimator(new DefaultItemAnimator());
    mRecycler.setHasFixedSize(true);

    mAdapter = new CardsAdapter(this);
    mRecycler.setAdapter(mAdapter);

    getSupportLoaderManager().initLoader(mAdapter.getLoaderId(), null, mAdapter);
  }

  @Override
  public boolean onCreateOptionsMenu(Menu menu) {
    getMenuInflater().inflate(R.menu.menu_main, menu);
    return true;
  }
}
