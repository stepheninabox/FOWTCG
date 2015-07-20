package com.inabox.stephen.fowtcg;

import android.app.Activity;
import android.graphics.drawable.Drawable;
import android.os.Parcelable;
import android.support.v7.app.ActionBarActivity;
import android.os.Bundle;
import android.view.Menu;
import android.view.MenuItem;
import android.view.View;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.ListView;

import java.io.InputStream;
import java.util.List;
import java.util.Random;


public class MainActivity extends Activity {

    private ListView listView;
    private ItemArrayAdapter itemArrayAdapter;
    private ImageView appImageView;
    private Button appButton;
    private Drawable drawable;
    private Drawable [] drawables = null;
    private Random random;



    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        appImageView = (ImageView) findViewById(R.id.imageView);
        appButton = (Button) findViewById(R.id.button);

        drawables = new Drawable[]{
                getResources().getDrawable(R.drawable.e_cmf_046),
                getResources().getDrawable(R.drawable.e_cmf_067)
        };

        appButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

                random = new Random();
                int randomNumber = random.nextInt(drawables.length);
                drawable = drawables[randomNumber];
                appImageView.setImageDrawable(drawable);
            }
        });

        listView = (ListView) findViewById(R.id.list_view);
        itemArrayAdapter = new ItemArrayAdapter(getApplicationContext(),R.layout.single_list_item);

        Parcelable state = listView.onSaveInstanceState();
        listView.setAdapter(itemArrayAdapter);
        listView.onRestoreInstanceState(state);

        InputStream inputStream = getResources().openRawResource(R.raw.playingcard);
        CSVReader csv = new CSVReader(inputStream);
        List<String[]> playerList = csv.read();

        for (String [] scoreData : playerList) {
            itemArrayAdapter.add(scoreData);
        }


    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu){
        getMenuInflater().inflate(R.menu.menu_main, menu);
        return true;

    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item){
        int id = item.getItemId();
        if (id == R.id.action_settings){
            return true;
        }
        return super.onOptionsItemSelected(item);
    }
}
