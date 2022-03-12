package com.shuaijun.brs.ui.classes

import android.view.ViewGroup
import androidx.recyclerview.widget.RecyclerView
import androidx.viewbinding.ViewBinding

class Adapter<V : ViewBinding, D : Any>(
    var viewHolder: (parent: ViewGroup, viewType: Int) -> V,
    var onBinderViewHolder: (binding: V, position: D) -> Unit,
    var list: MutableList<D>
) :
    RecyclerView.Adapter<VH<V>>() {

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): VH<V> =
        VH(viewHolder(parent, viewType))


    override fun onBindViewHolder(holder: VH<V>, position: Int) {
        onBinderViewHolder(holder.binding, list[position])
    }

    override fun getItemCount(): Int = list.size
}