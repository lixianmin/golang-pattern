package main

import (
	"github.com/lixianmin/dbi"
)

/********************************************************************
created:    2022-03-19
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

func updateMySQL(db *dbi.DB, ctx *dbi.Context, userId int) error {
	var tx, _ = db.BeginTx(ctx, nil)
	defer tx.Rollback()

	var count int
	if err := tx.SelectContext(ctx, &count, "select count from t where userId = ?", userId); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, "update t set count= ? where userId=?", count+1, userId); err != nil {
		return err
	}

	return tx.Commit()
}
