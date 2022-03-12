package com.shuaijun.brs.ui.classes

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import androidx.recyclerview.widget.LinearLayoutManager
import com.shuaijun.brs.databinding.FragmentClassesCatBinding
import com.shuaijun.brs.databinding.ItemFragmentCalssesCatBinding
import com.shuaijun.brs.retrofit.RF
import rx.android.schedulers.AndroidSchedulers
import rx.schedulers.Schedulers


class ClassesCatFragment : Fragment() {

    private lateinit var binding: FragmentClassesCatBinding
    private lateinit var adapter: Adapter<ItemFragmentCalssesCatBinding, Classes.ClassesDTO>

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View =
        FragmentClassesCatBinding.inflate(inflater, container, false).apply { binding = this }.root

    override fun onActivityCreated(savedInstanceState: Bundle?) {
        super.onActivityCreated(savedInstanceState)

        Adapter<ItemFragmentCalssesCatBinding, Classes.ClassesDTO>(
            { parent, _ ->
                ItemFragmentCalssesCatBinding.inflate(
                    LayoutInflater.from(requireContext()),
                    parent,
                    false
                )
            }, { binding, item ->
                binding.label.text = item.toString()
            }, mutableListOf()
        ).apply {
            binding.recyclerview.adapter = this
            binding.recyclerview.layoutManager = LinearLayoutManager(requireContext())
            adapter = this
        }

        RF.getInstance().classesList("all")
            .subscribeOn(Schedulers.io())
            .observeOn(AndroidSchedulers.mainThread())
            .subscribe {
                adapter.list = it.classes
                adapter.notifyDataSetChanged()
            }
    }

    companion object {
        @JvmStatic
        fun newInstance() = ClassesCatFragment()
    }
}