package com.shuaijun.brs.ui.material

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import androidx.recyclerview.widget.LinearLayoutManager
import com.shuaijun.brs.databinding.FragmentMaterialCatBinding
import com.shuaijun.brs.databinding.ItemFragmentCalssesCatBinding
import com.shuaijun.brs.databinding.ItemFragmentMaterialCatBinding
import com.shuaijun.brs.retrofit.RF
import com.shuaijun.brs.ui.classes.Adapter
import com.shuaijun.brs.ui.classes.Classes
import rx.android.schedulers.AndroidSchedulers
import rx.schedulers.Schedulers

class MaterialCatFragment : Fragment() {

    private lateinit var binding: FragmentMaterialCatBinding
    private lateinit var adapter: Adapter<ItemFragmentMaterialCatBinding, Material.MaterialDTO>

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View =
        FragmentMaterialCatBinding.inflate(inflater, container, false).apply { binding = this }.root


    override fun onStart() {
        super.onStart()
        Adapter<ItemFragmentMaterialCatBinding, Material.MaterialDTO>(
            { parent, _ ->
                ItemFragmentMaterialCatBinding.inflate(
                    LayoutInflater.from(requireContext()),
                    parent,
                    false
                )
            }, { holder, item ->
                val binding = holder.binding
                binding.labelName.text = String.format("物料名称:%s", item.name)
                binding.labelMid.text = String.format("物料编号:%s", item.mid)
                binding.labelCid.text = String.format("类别编号:%s", item.cid)

                binding.delete.setOnClickListener {
                    RF.getInstance().materialDelete(item.mid)
                        .subscribeOn(Schedulers.io())
                        .observeOn(AndroidSchedulers.mainThread())
                        .subscribe()
                    binding.root.close(true)
                    adapter.getData().removeAt(holder.adapterPosition)
                    adapter.notifyItemRemoved(holder.adapterPosition)
                }

            }, mutableListOf()
        ).apply {
            binding.recyclerview.adapter = this
            binding.recyclerview.layoutManager =
                LinearLayoutManager(requireContext(), LinearLayoutManager.VERTICAL, false)
            adapter = this
        }

        RF.getInstance().materialList("all")
            .subscribeOn(Schedulers.io())
            .observeOn(AndroidSchedulers.mainThread())
            .subscribe {
                adapter.list = it.materials
                adapter.notifyItemRangeChanged(0, it.materials.size)
            }
    }

}