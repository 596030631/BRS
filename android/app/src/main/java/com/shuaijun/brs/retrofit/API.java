package com.shuaijun.brs.retrofit;

import com.shuaijun.brs.ui.classes.Classes;
import com.shuaijun.brs.ui.material.Material;

import retrofit2.http.GET;
import retrofit2.http.Query;
import rx.Observable;

public interface API {

    @GET("classes/add")
    Observable<Classes> classesAdd(@Query("cid") String cid, @Query("name") String name, @Query("pid") String pid);

    @GET("classes/list")
    Observable<Classes> classesList(@Query("pid") String pid);

    @GET("classes/delete")
    Observable<Classes> classesDelete(@Query("cid") String cid);

    @GET("material/add")
    Observable<Material> materialAdd(@Query("mid") String mid, @Query("cid") String cid, @Query("name") String name, @Query("icon") String icon);

    @GET("material/list")
    Observable<Material> materialList(@Query("cid") String cid);

    @GET("material/delete")
    Observable<Material> materialDelete(@Query("cid") String cid);

}
