package com.shuaijun.brs.ui.material

import android.app.AlertDialog
import android.content.Context
import android.os.Bundle
import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.view.inputmethod.InputMethodManager
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProvider
import com.google.android.material.snackbar.Snackbar
import com.shuaijun.brs.databinding.FragmentMaterialAddBinding
import com.shuaijun.brs.retrofit.RF
import rx.android.schedulers.AndroidSchedulers
import rx.schedulers.Schedulers

class MaterialAddFragment : Fragment() {

    companion object {
        fun newInstance() = MaterialAddFragment()
    }

    private lateinit var viewModel: MaterialAddViewModel
    private lateinit var binding: FragmentMaterialAddBinding
    private var classesDialog: AlertDialog? = null

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View =
        FragmentMaterialAddBinding.inflate(inflater, container, false).apply { binding = this }.root

    override fun onActivityCreated(savedInstanceState: Bundle?) {
        super.onActivityCreated(savedInstanceState)
        viewModel = ViewModelProvider(this).get(MaterialAddViewModel::class.java)
        binding.root.setOnClickListener {
            val imm: InputMethodManager =
                requireActivity().getSystemService(Context.INPUT_METHOD_SERVICE) as InputMethodManager
            imm.hideSoftInputFromWindow(binding.inputCid.windowToken, 0)
        }

        binding.btnSubmit.setOnClickListener {
            val imm: InputMethodManager =
                requireActivity().getSystemService(Context.INPUT_METHOD_SERVICE) as InputMethodManager
            imm.hideSoftInputFromWindow(binding.inputCid.windowToken, 0)
            val cid = binding.inputCid.text.toString()
            val name = binding.inputName.text.toString()
            val pid = binding.btnPid.text.toString().split('\t')[0]
            if (cid.length == 4 && pid.isNotEmpty()) {
                RF.getInstance().materialAdd(cid, pid, name, "icon")
                    .subscribeOn(Schedulers.io())
                    .observeOn(AndroidSchedulers.mainThread())
                    .subscribe {
                        Log.d("TAG", it.toString())
                        if (it.code == 10000) {
                            Snackbar.make(binding.root, "新增物料成功", Snackbar.LENGTH_SHORT).show()
                        } else if (it.code == 11011) {
                            Snackbar.make(binding.root, "物料编号重复", Snackbar.LENGTH_SHORT).show()
                        } else {
                            Snackbar.make(binding.root, it.msg, Snackbar.LENGTH_SHORT).show()
                        }
                    }
            } else {
                Snackbar.make(binding.root, "请完善信息", Snackbar.LENGTH_SHORT).show()
            }
        }
    }

    override fun onStart() {
        super.onStart()
        binding.btnPid.keyListener = null
        binding.btnPid.setOnClickListener {
            classesDialog?.show()
        }
        RF.getInstance().classesList("all")
            .map {
                val array = arrayListOf<String>()
                for (i in it.classes) {
                    array.add(i.cid + "\t" + i.name)
                }
                return@map array.toTypedArray()
            }
            .subscribeOn(Schedulers.io()).observeOn(AndroidSchedulers.mainThread())
            .subscribe { data ->
                AlertDialog.Builder(requireContext())
                    .setItems(data) { dialog, which ->
                        Log.d("TAG", "which=$which")
                        binding.btnPid.setText(data[which])
                        dialog?.dismiss()
                    }
                    .create()
                    .apply {
                        classesDialog = this
                    }
                    .show()
            }
    }
}