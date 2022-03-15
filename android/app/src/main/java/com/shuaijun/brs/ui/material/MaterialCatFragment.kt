package com.shuaijun.brs.ui.material

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import com.shuaijun.brs.databinding.FragmentMaterialCatBinding

class MaterialCatFragment : Fragment() {

    private lateinit var binding: FragmentMaterialCatBinding

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View =
        FragmentMaterialCatBinding.inflate(inflater, container, false).apply { binding = this }.root

}