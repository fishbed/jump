package main

import (
    "fmt"
    "runtime"
    math "math"

    "github.com/go-gl/gl/v2.1/gl"
    "github.com/go-gl/glfw/v3.2/glfw"
)

const windowWidth  = 800
const windowHeight = 800


func init() {
    runtime.LockOSThread()

    // glfw 初期化
    if err := glfw.Init(); err != nil {
        panic(err)
    }

}

func main() {
    defer glfw.Terminate()

    // アプリウィンドウ表示
    window, err := glfw.CreateWindow(windowWidth, windowHeight, "Game of jump", nil, nil) 
    if err != nil {
        panic(err)
    }
    window.MakeContextCurrent()

    // gl 初期化
    if err := gl.Init(); err != nil {
        panic(err)
    }
    fmt.Println("OpenGL version", gl.GoStr(gl.GetString(gl.VERSION)))

    setScene ()
    // 表示用のループ
    for !window.ShouldClose() {
        drawScene()
        window.SwapBuffers()
        glfw.PollEvents()
    }
}

/**
 * アプリウィンドウの設定
 */
func setupWindow ()  {
    glfw.WindowHint(glfw.Resizable, glfw.False)
    glfw.WindowHint(glfw.ContextVersionMajor, 4)
    glfw.WindowHint(glfw.ContextVersionMinor, 1)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}


/**
 * 画面表示
 */
func drawScene () {
        gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)


        // ステージ表示
        drawStage()

        // 自機表示
        key := getDownKey()
        drawMyselfObject (key)
        
        // 敵表示
        drawEnemy_1 ()
        drawEnemy_2 ()

        gl.Flush()

}

/**
 * キャラクタデータ用の構造体
 */
type character struct {
	weight float32 
	height float32
	x      float32
	y      float32
} 
/**
 * スクリーンの設定
 */
func setScene () {
    gl.ClearColor(0.175, 0.175, 0.175, 0.0)
}

/**
 * 画面オブジェクト
 */

 /**
  * 入力キーの取得
  */
 func getDownKey()  string {
     return "";
 }

 /**
  * 自機オブジェクトの表示
  */
 func drawMyselfObject (downKey string) {

     // 半径
     radius :=0.05
     // 中心座標
     var cx float64 
     var cy float64

    // 自機用の円を表示
    gl.Begin(gl.TRIANGLES)
    gl.Color3f(0.8,0,0)
    
    for th1 := 0.0; th1 < 360.0; th1++ {

        th2 := th1 + 10.0

        th1Radius := th1 / 180.0 * math.Pi
        th2Radius := th2 / 180.0 * math.Pi

        x1 := radius * math.Cos(th1Radius)
        y1 := radius * math.Sin(th1Radius)
        x2 := radius * math.Cos(th2Radius)
        y2 := radius * math.Sin(th2Radius)

        // 円の描画
        gl.Vertex2f( float32(cx),    float32(cy)    ) // 中心点
        gl.Vertex2f( float32(x1+cx), float32(y1+cy) ) // 第一の円周上点
        gl.Vertex2f( float32(x2+cx), float32(y2+cy) ) // 次の円周上点

    }
    gl.End()

 }

 /**
  * 敵オブジェクトの表示
  */
  func drawEnemy_1()  {
    // 敵の諸元設定
    var enemy character
    enemy.weight = 0.1
    enemy.height = 0.1
    enemy.x      = 0.1
    enemy.y      =-0.05
      
    gl.Begin(gl.QUADS)
    gl.Color3f(0.5,0.5,0.5)
    gl.Vertex2f(enemy.x               , enemy.y)
    gl.Vertex2f(enemy.x + enemy.weight, enemy.y)
    gl.Vertex2f(enemy.x + enemy.weight, enemy.y - enemy.height)
    gl.Vertex2f(enemy.x               , enemy.y - enemy.height)
    gl.End()
      
  }
  func drawEnemy_2()  {
    // 敵の諸元設定
    var enemy character
    enemy.weight = 0.1
    enemy.height = 0.1
    enemy.x      =-0.2
    enemy.y      =-0.05

    gl.Begin(gl.QUADS)
    gl.Color3f(0.75,0.75,0.75)
    gl.Vertex2f(enemy.x               , enemy.y)
    gl.Vertex2f(enemy.x + enemy.weight, enemy.y)
    gl.Vertex2f(enemy.x + enemy.weight, enemy.y - enemy.height)
    gl.Vertex2f(enemy.x               , enemy.y - enemy.height)
    gl.End()
      
  }

 /**
  * ステージ表示
  */
  func drawStage()  {
    
	// ステージの諸元設定
    var stage character
    stage.weight =  1.8
    stage.height =  1.8
    stage.x      = -0.9
    stage.y      =  0.9

    // ステージ表示
    gl.Begin(gl.QUADS)
    gl.Color3f(1.0,1.0,1.0)
    gl.Vertex2f(stage.x               , stage.y)
    gl.Vertex2f(stage.x + stage.weight, stage.y)
    gl.Color3f(0.5,0.5,0.5)
    gl.Vertex2f(stage.x + stage.weight, stage.y - stage.height)
    gl.Vertex2f(stage.x               , stage.y - stage.height)
    gl.End()
  }

 /**
  * 場外判定
  */
  func checkOnStage() bool  {
    return true
  }
