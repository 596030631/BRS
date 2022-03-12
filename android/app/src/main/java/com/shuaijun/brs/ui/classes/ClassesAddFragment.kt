package com.shuaijun.brs.ui.classes

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

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View =
        FragmentClassesAddBinding.inflate(inflater, container, false).apply { binding = this }.root


    override fun onActivityCreated(savedInstanceState: Bundle?) {
        super.onActivityCreated(savedInstanceState)
        viewModel = ViewModelProvider(this).get(ClassesAddViewModel::class.java)

        binding.btnSubmit.setOnClickListener { v ->
            val imm: InputMethodManager =
                requireActivity().getSystemService(Context.INPUT_METHOD_SERVICE) as InputMethodManager
            imm.hideSoftInputFromWindow(binding.inputCid.windowToken, 0)
            val cid = binding.inputCid.text.toString()
            val name = binding.inputName.text.toString()
            val pid = binding.btnPid.text.toString()
            if (cid.length > 6 && name.length > 3) {
                RF.getInstance().classesAdd(cid, name, pid)
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
}