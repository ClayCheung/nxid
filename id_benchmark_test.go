package nxid

import (
	"testing"
	//"github.com/google/uuid"
	//"github.com/rs/xid"
	//guuid "github.com/satori/go.uuid"
)

func BenchmarkNew(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = New()
		}
	})
}

func BenchmarkFromString(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = FromString("2vmltfb8ksvggo7kgri2gg9dp4")
		}
	})
}

func BenchmarkNXID(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = New().String()
		}
	})
}

//func BenchmarkXID(b *testing.B) {
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			_ = xid.New().String()
//		}
//	})
//}

//func BenchmarkGoUUIDv1(b *testing.B) {
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			_ = guuid.NewV1().String()
//		}
//	})
//}
//
//func BenchmarkUUIDv4(b *testing.B) {
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			_ = uuid.New().String()
//		}
//	})
//}
//
//func BenchmarkUUIDv6(b *testing.B) {
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			u, _ := uuid.NewV6()
//			_ = u.String()
//		}
//	})
//}
//func BenchmarkUUIDv7(b *testing.B) {
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			u, _ := uuid.NewV7()
//			_ = u.String()
//		}
//	})
//}
