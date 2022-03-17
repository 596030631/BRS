package com.shuaijun.brs.retrofit;

import com.shuaijun.brs.ui.classes.Classes;
import com.shuaijun.brs.ui.login.LoggedInUser;
import com.shuaijun.brs.ui.material.Material;

import retrofit2.Call;
import retrofit2.http.GET;
import retrofit2.http.Query;
import rx.Observable;

public interface API {

    @GET("user/login")
    Call<LoggedInUser> userLogin(@Query("user_id") String userName, @Query("passwd") String passwd);

    @GET("classes/add")
    Observable<Classes> classesAdd(@Query("cid") String classesId, @Query("name") String name, @Query("pid") String parentClassesId);

    @GET("classes/list")
    Observable<Classes> classesList(@Query("pid") String parentClassesId);

    @GET("classes/delete")
    Observable<Classes> classesDelete(@Query("cid") String classesId);

    @GET("material/add")
    Observable<Material> materialAdd(@Query("mid") String materialId, @Query("cid") String classesId, @Query("name") String name, @Query("icon") String icon);

    @GET("material/list")
    Observable<Material> materialList(@Query("cid") String classesId);

    @GET("material/delete")
    Observable<Material> materialDelete(@Query("mid") String materialId);

}
