package banana

import "errors"

var (
	// repo
	RepoNotUpdated = errors.New("Cập nhật thông tin Repo thất bại")
	RepoNotFound   = errors.New("Repo không tồn tại")
	RepoConflict   = errors.New("Repo đã tồn tại")
	RepoInsertFail = errors.New("Thêm Repo thất bại")

	// bookmark
	BookmarkNotFound   = errors.New("Bookmark không tồn tại")
	BookmarkConflict   = errors.New("Bookmark đã tồn tại")
	BookmarkInsertFail = errors.New("Thêm Bookmark thất bại")
	BookmarkDeleteFail = errors.New("Xóa Bookmark thất bại")

	ErrorSql = errors.New("Lỗi SQL")
)
