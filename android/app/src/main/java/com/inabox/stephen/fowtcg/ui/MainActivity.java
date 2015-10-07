package com.inabox.stephen.fowtcg.ui;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.DefaultItemAnimator;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.text.Editable;
import android.text.TextWatcher;
import android.view.Menu;
import android.widget.EditText;
import butterknife.ButterKnife;
import butterknife.InjectView;
import com.inabox.stephen.fowtcg.R;
import go.game.Game;


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

    getSupportLoaderManager().initLoader(mAdapter.getLoaderId(), new Bundle(), mAdapter);
  }

  @Override
  public boolean onCreateOptionsMenu(Menu menu) {
    getMenuInflater().inflate(R.menu.menu_main, menu);

    EditText search = (EditText) menu.findItem(R.id.action_search)
        .getActionView()
        .findViewById(R.id.search);

    search.addTextChangedListener(new TextWatcher() {
      @Override
      public void beforeTextChanged(CharSequence s, int start, int count, int after) {
      }

      @Override
      public void onTextChanged(CharSequence s, int start, int before, int count) {
      }

      @Override
      public void afterTextChanged(Editable s) {
        String title = s.toString();
        try {
          mAdapter.mCards = Game.FindByTitle(title);
          mAdapter.notifyDataSetChanged();
        } catch (Exception e) {
          e.printStackTrace();
        }

//        Bundle args = new Bundle();
//        args.putString("filter", s.toString());
//        mAdapter.hintRemoveInsert();
//        getSupportLoaderManager().restartLoader(mAdapter.getLoaderId(), args, mAdapter);
      }
    });

    return super.onCreateOptionsMenu(menu);
  }
}
