package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	positions := []string{
		"主人公の現在",
		"主人公の近い未来",
		"主人公の過去",
		"援助者",
		"敵対者",
		"結末(目的)",
	}

	attributes := []string{
		"意思",
		"清楚",
		"理性",
		"秩序",
		"知恵",
		"制約",
		"善良",
		"節度",
		"至誠",
		"生命",
		"寛容",
		"開放",
		"調和",
		"創造",
		"信頼",
		"公式",
		"変化",
		"結合",
		"厳格",
		"勇気",
		"幸運",
		"庇護",
		"治癒",
		"慈愛",
	}

	// 各立ち位置に対してランダムな属性を割り当てる
	for _, position := range positions {
		selectedAttribute := attributes[rand.Intn(len(attributes))]

		// 50％の確率で"逆"を前に追加
		if rand.Float32() < 0.5 {
			selectedAttribute = "逆" + selectedAttribute
		}

		fmt.Printf("%s: %s\n", position, selectedAttribute)
	}
}
