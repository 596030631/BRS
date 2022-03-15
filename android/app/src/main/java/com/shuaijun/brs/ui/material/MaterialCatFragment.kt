package com.shuaijun.brs.ui.material

import android.content.Context
import android.os.Bundle
import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.view.inputmethod.InputMethodManager
import androidx.fragment.app.Fragment
import com.google.android.material.snackbar.Snackbar
import com.shuaijun.brs.databinding.FragmentMaterialCatBinding
import com.shuaijun.brs.retrofit.RF
import rx.android.schedulers.AndroidSchedulers
import rx.schedulers.Schedulers

class MaterialCatFragment : Fragment() {

    private lateinit var binding: FragmentMaterialCatBinding

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View =
        FragmentMaterialCatBinding.inflate(inflater, container, false).apply { binding = this }.root




}