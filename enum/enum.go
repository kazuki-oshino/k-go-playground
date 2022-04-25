//go:generate stringer -type=GameType

package enum

// GoではEnumは存在しない
// typeとconstをコラボして書いたりする
// 文字列にしたいときはstringerパッケージ等で生成すると良い
type GameType int

// iotaでシーケンスを振ったり、
// 固定的な番号を書いたり、
// bit演算できるように書いたり、
// 様々な方法を駆使する
const (
	PS4 GameType = iota + 1
	PS5
	Wii
	WiiU
	NintendoSwitch
)
