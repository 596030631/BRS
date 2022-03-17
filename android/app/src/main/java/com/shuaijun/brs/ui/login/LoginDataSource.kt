package com.shuaijun.brs.ui.login

import com.shuaijun.brs.retrofit.RF
import java.io.IOException
import java.net.SocketException

/**
 * Class that handles authentication w/ login credentials and retrieves user information.
 */
class LoginDataSource {

    fun login(username: String, password: String): Result<LoggedInUser> {
        try {
            val response = RF.getInstance().userLogin(username, password).execute()
            return if (response.isSuccessful) {
                Result.Success(response.body())
            } else {
                Result.Error(SocketException("连接失败"))
            }
        } catch (e: Throwable) {
            return Result.Error(IOException("Error logging in", e))
        }
    }

    fun logout() {
        // TODO: revoke authentication
    }
}