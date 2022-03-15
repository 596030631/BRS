package com.shuaijun.brs.ui.home

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProvider
import androidx.navigation.Navigation
import com.shuaijun.brs.R
import com.shuaijun.brs.databinding.FragmentHomeBinding

class HomeFragment : Fragment() {

    private lateinit var homeViewModel: HomeViewModel
    private var _binding: FragmentHomeBinding? = null

    // This property is only valid between onCreateView and
    // onDestroyView.
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        homeViewModel =
            ViewModelProvider(this).get(HomeViewModel::class.java)

        _binding = FragmentHomeBinding.inflate(inflater, container, false)

        val root: View = binding.root

        binding.btnClassesAdd.setOnClickListener {
            Navigation.findNavController(it).navigate(R.id.action_nav_home_to_classesAddFragment)
        }

        binding.btnClassesCat.setOnClickListener {
            Navigation.findNavController(it).navigate(R.id.action_nav_home_to_classesCatFragment)
        }

        binding.btnMaterialAdd.setOnClickListener {
            Navigation.findNavController(it).navigate(R.id.action_nav_home_to_materialAddFragment)
        }

        binding.btnMaterialCat.setOnClickListener {
            Navigation.findNavController(it).navigate(R.id.action_nav_home_to_nav_material_cat)
        }

        homeViewModel.text.observe(viewLifecycleOwner, {

        })
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }
}