package com.shuaijun.brs.retrofit;

import com.shuaijun.brs.ui.classes.Classes;

import java.util.List;

import retrofit2.http.GET;
import retrofit2.http.Query;
import rx.Observable;

public interface API {

    @GET("classes/add")
    Observable<Classes> classesAdd(@Query("cid") String cid, @Query("name") String name, @Query("pid") String pid);

    @GET("classes/list")
    Observable<Classes> classesList(@Query("pid") String pid);

}
