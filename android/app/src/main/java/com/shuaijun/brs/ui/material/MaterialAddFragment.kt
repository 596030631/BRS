package com.shuaijun.brs.ui.material

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProvider
import com.shuaijun.brs.databinding.FragmentMaterialAddBinding

class MaterialAddFragment : Fragment() {

    companion object {
        fun newInstance() = MaterialAddFragment()
    }

    private lateinit var viewModel: MaterialAddViewModel
    private lateinit var binding: FragmentMaterialAddBinding

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View =
        FragmentMaterialAddBinding.inflate(inflater, container, false).apply { binding = this }.root


    override fun onActivityCreated(savedInstanceState: Bundle?) {
        super.onActivityCreated(savedInstanceState)
        viewModel = ViewModelProvider(this).get(MaterialAddViewModel::class.java)
    }
}