#!/bin/bash
set -e

# 打印開始消息
echo "開始執行資料庫初始化腳本..."
echo "等待 PostgreSQL 服務器啟動..."
sleep 10  # 等待 10 秒
# 嘗試連接到 PostgreSQL 服務器並創建資料庫
echo "正在連接到 PostgreSQL 服務器..."
psql -h postgres -U royce -d stock_info -c "CREATE DATABASE stock_info_distributor;";
echo "資料庫 stock_info_distributor 創建完成（如果之前不存在）。"

psql -h postgres -U royce -d stock_info -c "CREATE DATABASE stock_info_scheduler;";
echo "資料庫 stock_info_scheduler 創建完成（如果之前不存在）。"

# 打印完成消息
echo "資料庫初始化腳本執行完畢。"