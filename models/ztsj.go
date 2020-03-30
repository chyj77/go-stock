package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Ztsj struct {
	Recid   int64  `json:"recid,string" orm:"column(recid);pk;auto;"`
	Rq      string `json:"rq" orm:"column(rq)"`
	Mrztgs  string `json:"mrztgs" orm:"column(mrztgs)"`
	Fyzbgs  string `json:"fyzbgs" orm:"column(fyzbgs)"`
	Ztzdgn  string `json:"ztzdgn" orm:"column(ztzdgn)"`
	Ztzdgs  string `json:"ztzdgs" orm:"column(ztzdgs)"`
	Dbcrgkl string `json:"dbcrgkl" orm:"column(dbcrgkl)"`
	Spcgl   string `json:"spcgl" orm:"column(spcgl)"`
	Bzsl    string `json:"bzsl" orm:"column(bzsl)"`
	Bzl     string `json:"bzl" orm:"column(bzl)"`
}

type ZtsjResponse struct {
	PageIndex   int    `json:"pageIndex"`
	PageSize    int    `json:"pageSize"`
	TotalCounts int64  `json:"totalCounts"`
	Datas       []Ztsj `json:"data"`
}

func QueryZtsj(pageIndex, pageSize int) (int64, []Ztsj) {
	//ormer
	var myOrmer orm.Ormer
	myOrmer = orm.NewOrm()
	var ztsjs []Ztsj
	var ztsjSql = "SELECT ztsj.rq AS rq,ztsj.mrztgs AS mrztgs,ztsj.fyzbgs AS fyzbgs," +
		" ztsj.ztzdgn AS ztzdgn,ztsj.ztzdgs AS ztzdgs," +
		" ztsj.dbcrgkl AS dbcrgkl," +
		" ztsj.spcgl AS spcgl,ztsj.recid as recid," +
		" ztsj.bzsl AS bzsl," +
		" ztsj.bzl AS bzl" +
		" FROM ztsj" +
		" order by rq DESC " +
		" limit ?,? "
	myOrmer.Raw(ztsjSql, pageIndex, pageSize).QueryRows(&ztsjs)
	ztsjCounts, err := myOrmer.QueryTable("ztsj").Count()

	if err == nil {
		fmt.Println("total countes = ", ztsjCounts)
	}
	return ztsjCounts, ztsjs
}

func SaveZtsj(z *Ztsj) int64 {
	//ormer
	var myOrmer orm.Ormer
	myOrmer = orm.NewOrm()
	recid := z.Recid
	var err error
	var num int64
	if recid == 0 {
		num, err = myOrmer.Insert(z)
	} else {
		num, err = myOrmer.Update(z)
	}
	fmt.Println(num)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return num
}

func DeleteZtsj(recid int64) int64 {
	//ormer
	var myOrmer orm.Ormer
	myOrmer = orm.NewOrm()
	var ztsj Ztsj
	ztsj.Recid = recid
	num, err := myOrmer.Delete(&ztsj)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return num
}
