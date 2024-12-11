package broker

import "fmt"

/**
*
* kafka
* <p>
* kafka file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is shared under a collaborative license.
* Contributions, suggestions, and improvements are welcome!
* Feel free to fork, modify, and submit pull requests under the terms of the repository's license.
* Please ensure proper attribution to the original author(s) and maintain this notice in derivative works.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 7/12/2024
*
 */

type KafkaMq struct {
}

func newKafka() *KafkaMq {
	return &KafkaMq{}
}

func (k *KafkaMq) Process() {
	fmt.Println("init kafka process")
}
