package com.shuaijun.brs.ui.login

import com.google.gson.annotations.SerializedName

/**
 * Data class that captures user information for logged in users retrieved from LoginRepository
 */
data class LoggedInUser(
    @SerializedName("user_id")
    val userId: String,
    @SerializedName("user_name")
    val displayName: String,
    val passwd: String,
    @SerializedName("user_sex")
    val userSex: String,
    @SerializedName("user_age")
    val userAge: String,
    @SerializedName("user_level")
    val userLevel: Int,
    @SerializedName("user_icon")
    val userIcon: String,
)