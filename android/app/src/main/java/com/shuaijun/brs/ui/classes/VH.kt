package com.shuaijun.brs.ui.classes

import androidx.recyclerview.widget.RecyclerView
import androidx.viewbinding.ViewBinding

data class VH<T : ViewBinding>(var binding: T) : RecyclerView.ViewHolder(binding.root)

