package com.shuaijun.brs.retrofit;

import okhttp3.OkHttpClient;
import retrofit2.Retrofit;
import retrofit2.adapter.rxjava.RxJavaCallAdapterFactory;
import retrofit2.converter.gson.GsonConverterFactory;

public class RF {
    private static final OkHttpClient okHttpClient = new OkHttpClient.Builder()
            .addInterceptor(new LoggerInterceptor())
            .build();
    private static final Retrofit retrofit = new Retrofit.Builder()
            .baseUrl("http://inlets.fun:8080/")
            .addConverterFactory(GsonConverterFactory.create())
            .addCallAdapterFactory(RxJavaCallAdapterFactory.create())
            .client(okHttpClient)
            .build();
    public static API getInstance() {
        return retrofit.create(API.class);
    }

}
