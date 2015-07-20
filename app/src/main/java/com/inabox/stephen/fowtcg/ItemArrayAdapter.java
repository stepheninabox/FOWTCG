package com.inabox.stephen.fowtcg;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.TextView;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Stephen on 7/13/2015.
 */
public class ItemArrayAdapter extends ArrayAdapter<String[]>{

    private List<String[]> playerList = new ArrayList<String[]>();

    static class ItemViewHolder {
        TextView name;
        TextView set;
    }

    public ItemArrayAdapter(Context context, int resource) {
        super(context, resource);
    }

    public void add(String[] object) {
        playerList.add(object);
        super.add(object);
    }

    @Override
    public int getCount(){
        return this.playerList.size();
    }

    @Override
    public String[] getItem(int position) {
        return this.playerList.get(position);
    }

    @Override
    public View getView(int position, View convertView, ViewGroup parent){
        View row = convertView;
        ItemViewHolder viewHolder;
        if(row== null){
            LayoutInflater inflater = (LayoutInflater) this.getContext().getSystemService(Context.LAYOUT_INFLATER_SERVICE);
            row = inflater.inflate(R.layout.single_list_item, parent, false);
            viewHolder = new ItemViewHolder();
            viewHolder.name = (TextView) row.findViewById(R.id.name);
            viewHolder.set = (TextView) row.findViewById(R.id.set);
            row.setTag(viewHolder);

        } else {
            viewHolder = (ItemViewHolder) row.getTag();
        }

        String[] stat = getItem(position);
        viewHolder.name.setText(stat[0]);
        viewHolder.set.setText(stat[1]);
        return row;
    }
}
