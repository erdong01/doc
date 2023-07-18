package pipeline

import (
	"fmt"
	"sync"
	"testing"
)

// Pipeline 模式也成为流水线模式，模拟现实中的流水线生成。我们以组装手机为例，假设只有三道工序：零件采购、组装、打包成品：
// 零件采购（工序1）-》组装（工序2）-》打包（工序3）
func TestXxx(t *testing.T) {
	ch := buy(10)
	zgch1 := build(ch)
	zgch2 := build(ch)
	zgch3 := build(ch)
	zgch := merge(zgch1, zgch2, zgch3)
	res := pack(zgch)
	for v := range res {
		fmt.Println(v)
	}
}

func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- fmt.Sprint("工序1：采购", i)
		}
	}()
	return out
}

func build(ch <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for v := range ch {
			out <- v + "工序2：组装"
		}
	}()
	return out
}

func pack(ch <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for v := range ch {
			out <- v + "工序2：打包"
		}
	}()
	return out
}

func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	wg.Add(len(ins))
	p := func(ch <-chan string) {
		defer wg.Done()
		for v := range ch {
			out <- v
		}
	}
	for _, v := range ins {
		go p(v)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// 题目数据
var QuestionData = map[int]Question{
	1001: {Ans: "A", Score: 10},
	1002: {Ans: "D", Score: 10},
	2001: {Ans: "C", Score: 10},
	2002: {Ans: "B", Score: 10},
	3001: {Ans: "A", Score: 10},
	3002: {Ans: "D", Score: 10},
	4001: {Ans: "A", Score: 10},
	4002: {Ans: "C", Score: 10},
	5001: {Ans: "D", Score: 10},
	5002: {Ans: "B", Score: 10},
}

type Question struct {
	No         int     // 题目编号
	Type       int     // 题型
	Ans        string  // 题目答案
	StudentAns string  // 学生答案
	Score      float32 // 分数
	IsRgiht    bool    // 对/错
}

func TestExam(t *testing.T) {
	//试卷
	var exam = map[int][]Question{
		1: {{No: 1001, StudentAns: "A", Type: 1}, {No: 1002, StudentAns: "D"}},
		2: {{No: 2001, StudentAns: "C", Type: 2}, {No: 2002, StudentAns: "B"}},
		3: {{No: 3001, StudentAns: "A", Type: 3}, {No: 3002, StudentAns: "D"}},
		4: {{No: 4001, StudentAns: "A", Type: 4}, {No: 4002, StudentAns: "C"}},
		5: {{No: 5001, StudentAns: "D", Type: 5}, {No: 5002, StudentAns: "B"}},
	}
	//开始阅卷 Start
	q := FindQuestion(exam)
	qRes1 := Scoring(q)
	qRes2 := Scoring(q)
	qRes3 := Scoring(q)
	_, score := report(qRes1, qRes2, qRes3)
	//开始阅卷 End
	fmt.Println("得分:", score)
}

// 查询试卷题目信息
func FindQuestion(exam map[int][]Question) <-chan []Question {
	out := make(chan []Question)
	go func() {
		defer close(out)
		for key := range exam {
			for k := range exam[key] {
				exam[key][k].Ans = QuestionData[exam[key][k].No].Ans
			}
			out <- exam[key]
		}
	}()
	return out
}

// 题目阅卷
func Scoring(in <-chan []Question) <-chan []Question {
	out := make(chan []Question)
	go func() {
		defer close(out)
		for c := range in {
			for k, question := range c {
				//处理学生题目数据
				if question.StudentAns == question.Ans {
					c[k].IsRgiht = true
					c[k].Score = QuestionData[question.No].Score
				} else {
					c[k].IsRgiht = false
				}
			}
			//其他操作。。。
			out <- c
		}
	}()
	return out
}

// 生成试卷报告
func report(ch ...<-chan []Question) (m map[int][]Question, score float32) {
	var wg sync.WaitGroup
	wg.Add(len(ch))
	m = make(map[int][]Question)
	p := func(c <-chan []Question) {
		defer wg.Done()
		for questions := range c {
			m[questions[0].Type] = questions
			for _, q := range questions {
				score += q.Score
			}
		}
	}
	for _, c := range ch {
		go p(c)
	}
	wg.Wait()
	return
}
