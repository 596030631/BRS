package com.shuaijun.brs.ui.classes

import android.annotation.SuppressLint
import android.os.Bundle
import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import androidx.recyclerview.widget.LinearLayoutManager
import com.shuaijun.brs.databinding.FragmentClassesCatBinding
import com.shuaijun.brs.databinding.ItemFragmentCalssesCat2Binding
import com.shuaijun.brs.retrofit.RF
import rx.android.schedulers.AndroidSchedulers
import rx.schedulers.Schedulers


class ClassesCatFragment : Fragment() {

    private lateinit var binding: FragmentClassesCatBinding
    private lateinit var adapter: Adapter<ItemFragmentCalssesCat2Binding, Classes.ClassesDTO>

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View =
        FragmentClassesCatBinding.inflate(inflater, container, false).apply { binding = this }.root

    override fun onActivityCreated(savedInstanceState: Bundle?) {
        super.onActivityCreated(savedInstanceState)

        Adapter<ItemFragmentCalssesCat2Binding, Classes.ClassesDTO>(
            { parent, _ ->
                ItemFragmentCalssesCat2Binding.inflate(
                    LayoutInflater.from(requireContext()),
                    parent,
                    false
                )
            }, { holder, item ->
                val binding = holder.binding
                binding.labelName.text = "类别名称:" + item.name
                binding.labelCid.text = "类别编号:" + item.cid
                binding.labelPid.text = "上级类别:" + item.pid

                binding.labelName.setOnClickListener {
                    Log.d("TAG", "CCCCC")
                }

                binding.delete.setOnClickListener {
                    RF.getInstance().classesDelete(item.cid)
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

        RF.getInstance().classesList("all")
            .subscribeOn(Schedulers.io())
            .observeOn(AndroidSchedulers.mainThread())
            .subscribe {
                adapter.list = it.classes
                adapter.notifyItemRangeChanged(0, it.classes.size)
            }
    }

    companion object {
        @JvmStatic
        fun newInstance() = ClassesCatFragment()
    }
}