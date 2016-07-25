/**
 * test ファイル
 * 実行コマンド ： go test jump.go jump_test.go
 */
package jump

import (
    "testing"
)

func TestChecOnStage(t *testing.T)  {
    onStage := checkOnStage()
    if !onStage {
        t.Error ("it outof stage!")
    }
}