package main

import (
    "encoding/json"
    "fmt"
    "os"
    "math"
)

type Data struct {
    Max      int     `json:"max"`
    Features int     `json:"features"`
    Points   []Point `json:"points"`
}

type Point struct {
    Label    string `json:"label"`
    Position []int  `json:"position"`
}

func main() {
    /**
    **
    ** Read Data
    **
    **/
    // Open the JSON file
    file, err := os.Open("data.json")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Decode the JSON data into a struct
    var data Data
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&data); err != nil {
        fmt.Println("Error:", err)
        return
    }

    /**
    **
    ** Random Inital Values
    **
    **/
    w := [3]float64{0, 1, 0.5}
    n := 0.2

    /**
    **
    ** iterate
    **
    **/
    j := 0
    for i := 0; i < data.Max; i++ {
        // iterate
        point := iterate(&data, &j)

        // check
        result := check(&point, w)
        if result {
            continue
        }

        // modify
        modify(&point, &w, &n)
    }
}

func iterate(data *Data, j *int) Point {
    if *j >= len(data.Points) {
        *j = 0
    }

    point := data.Points[*j]

    *j++
    return point
}

func check(p *Point, w [3]float64) bool {
    // s = 0
    a := ((w[0] * float64(p.Position[0])) + (w[1] * float64(p.Position[1]))) - 0

    if a > 0 {
        return true
    }

    return false
}

func modify(p *Point, w *[3]float64, n *float64) {
    var d = 1.0
    if p.Label == "b" {
        d = -1.0
    }

    w[0] = math.Round((w[0] + (*n * d)) * 10) / 10
    w[1] = math.Round((w[1] + (*n * d * float64(p.Position[0]))) * 10) / 10
    w[2] = math.Round((w[2] + (*n * d * float64(p.Position[1]))) * 10) / 10

    fmt.Printf("W: [%.1f, %.1f, %.1f]\n", w[0], w[1], w[2])
}

