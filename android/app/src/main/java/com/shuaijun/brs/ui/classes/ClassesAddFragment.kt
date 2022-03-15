package com.shuaijun.brs.ui.classes

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
import com.shuaijun.brs.databinding.FragmentClassesAddBinding
import com.shuaijun.brs.retrofit.RF
import rx.android.schedulers.AndroidSchedulers
import rx.schedulers.Schedulers

class ClassesAddFragment : Fragment() {

    companion object {
        fun newInstance() = ClassesAddFragment()
    }

    private lateinit var viewModel: ClassesAddViewModel
    private lateinit var binding: FragmentClassesAddBinding
    private var classesDialog: AlertDialog? = null

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View =
        FragmentClassesAddBinding.inflate(inflater, container, false).apply { binding = this }.root

    override fun onActivityCreated(savedInstanceState: Bundle?) {
        super.onActivityCreated(savedInstanceState)
        viewModel = ViewModelProvider(this).get(ClassesAddViewModel::class.java)

        RF.getInstance().classesList("all")
            .map {
                val array = arrayListOf<String>()
                for (i in it.classes) {
                    array.add(i.cid + "\t" + i.name)
                }
                return@map array.toTypedArray()
            }
            .subscribeOn(Schedulers.io()).observeOn(AndroidSchedulers.mainThread())
            .subscribe {data->
                AlertDialog.Builder(requireContext())
                    .setItems(data) { dialog, which ->
                        Log.d("TAG", "which=$which")
//                        binding.btnPid.text = data[which].split('\t')[0]
                        dialog?.dismiss()
                    }
                    .create()
                    .apply {
                        classesDialog = this
                    }
                    .show()
            }
//        binding.btnPid.setOnClickListener {
//            classesDialog?.show()
//        }

        binding.btnSubmit.setOnClickListener { v ->
            val imm: InputMethodManager =
                requireActivity().getSystemService(Context.INPUT_METHOD_SERVICE) as InputMethodManager
            imm.hideSoftInputFromWindow(binding.inputCid.windowToken, 0)
            val cid = binding.inputCid.text.toString()
            val name = binding.inputName.text.toString()
//            val pid = binding.btnPid.text.toString()
            if (cid.length == 3) {
                RF.getInstance().classesAdd(cid, name, "000")
                    .subscribeOn(Schedulers.io())
                    .observeOn(AndroidSchedulers.mainThread())
                    .subscribe {
                        Log.d("TAG", it.toString())
                        if (it.code == 10000) {
                            Snackbar.make(v, "新增成功", Snackbar.LENGTH_SHORT).show()
                        } else if (it.code == 11011) {
                            Snackbar.make(v, "类别已存在", Snackbar.LENGTH_SHORT).show()
                        } else {
                            Snackbar.make(v, it.msg, Snackbar.LENGTH_SHORT).show()
                        }
                    }
            } else {
                Snackbar.make(v, "请完善表单", Snackbar.LENGTH_SHORT).show()
            }
        }
    }

    override fun onPause() {
        super.onPause()
        val imm: InputMethodManager =
            requireActivity().getSystemService(Context.INPUT_METHOD_SERVICE) as InputMethodManager
        imm.hideSoftInputFromWindow(binding.inputCid.windowToken, 0)
    }
}