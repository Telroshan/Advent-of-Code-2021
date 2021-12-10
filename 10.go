package main

import "fmt"

func main() {
	inputRows := []string{
		"<[<<({{<{<[[([()][()()])<{[]<>}{<>}>><<<[]{}>>{<<><>>({}{})}>][<<<{}<>>([])>{<{}><()>}>]><(<{(()[]){[](",
		"<{([({{[(<[({({}{})[()<>]}{{<>{}}[(){}]})]>({(<<<><>>{{}()}>)<<<[]><{}<>>><[{}[]]<<>[]>>>}))[{(<{([]",
		"([<([({(<([<({()()}(<>[]))[<[]()>{<>{}}]>]<{{{{}{}}{()()})({[]}[<>{}])}>)>)}){<<{{(<<<()<>><<><>>>",
		"[{{[<<[{([[{{[<>{}]<()<>>}[(()<>)[{}<>]]}<<[{}[]]([]<>)>{{[]{}}{()()}}>]][[{<({}{})<[]{}>>",
		"<<{{<([(((([{<[]<>>{()<>}}(<<>[]>{(){}})]([{{}<>}<[]>](<[][]><{}()>)))))[({{{<[]{}>[()]}<<<><>>",
		"<(<[{(<[{{{<{[[]<>]{<>()}}><[[{}()][(){}]]>}}<[<[<()()>({}<>)]><[{[]()}([][])](<(){}>[<><>])>]{",
		"{[[<{{<<[[<<(({}<>){<><>})[[<>()][[]{})]>[[[{}[]]<<>>]{(<>[])}]>{{<<()[]>{<>{}}>[{[]<>}<[]{}>]",
		"({[((((([{[({[{}[]]<[]()>}[({}{})<<>[]>])<[(()[])<{}()>]{[[]<>]{[]()})>](([({}<>)(()[])][<[]{}>])",
		"{[[{[({(<<{<([{}<>])<({}<>)({}<>}>>[<[{}()]([]<>)>((<>())<[]{}>)]}<([(<>())[(){}]]{[{}<>]<[][]>}){<{<><",
		"(<{<<[<<([[([<<><>>]<<{}()><[]<>>>)]][{(({<>{}}(()())))(<[()()]<[]()>>)}[[<<<>()><<>{}>>([{}<>][{}[]])]{<{(",
		"<<[({{{{{[<([((){})([]())]<((){})({}[])>)[{[{}()]{[]<>}}]>([<{{}<>}><({}<>)({}[])>]<[{<>{}}",
		"(([[{{({{{[<<{[]()}({})>{{{}()}[[]<>]}>{<([])>[[[]{}]]}](([<[]{}>]<<[]<>><{}[]>>))}}}(([[[{[<><>",
		"<<[(<[{{<((<[<()>[<><>]][[<><>]([]())]>[({<>})<<()<>><{}[]>>])[[(<[]<>>([]())){{<>[]}(()())}]]){(<(<[]<>>",
		"({[<<[[{<[[{{[()()]{()[]}}}]<{([{}{}]{[]()})[[()[]]<<>>]}>]>}[({((<{<><>)>{{[]()}{{}[]}}){([{}[]]<{}{}>)<",
		"([{{<{[(<[<<[({}())({}())]>>[{[<<>[]>[<>()]]<(<>[]){()[]}>}(<<<>[]>[<>[]]>{<<>[]>{{}()}})]]{",
		"<[{{<[((<([[{<<><>>{()}}{<{}[]><[]()]}]])<{({<()>{()[]}}(<(){}><[]<>>)){[({}[]){<>[]}]}}<<([()()]<()(",
		"<{<{[(({(({({<[]>([])}<<()<>>{<>()}>){[[()[]]<{}()>]{({}[])[()<>]}}}<[<<{}()><[]{}>>({[][]}[()<>])}{<<[]",
		"<[{{[<({(<[<[[()<>]({}[])]([{}[]]([]))>]{[{{<>[]}[[]{}]}(({}<>)<()[]>)]<<[()[]]>({(){}}<{}()>)>}>)}<<[[[[<<>(",
		"{{<(<[((({(<[<[]{}>[(){}]](<{}<>>((){}))>){(({[]{}}<{}[]>){{<><>}[[]{}]}))}<([<(<>)>[<<>{}>[{}",
		"<([({[<[(({[(({}))(<{}()><<>{}>)]<[{{}[]}{{}<>}]{[{}()]{{}()}}>})([[<(<>())({}{})>[<[]<>>[",
		"{((<<<{<(<[[[<{}<>}]<(<><>)>]]<(([(){}]{[]()}))<({{}<>}<()<>>){<()<>>{(){}}}>>>{({({()[]}[<>",
		"([<[[((<(<[<([{}[]])>{<(()<>)((){})>{<<>>(()<>)}}][{[(<>[])(<>[])]}[[{{}[]}[()()]]<<[]<>>{(){}}>]]><({(",
		"[{<([[[{([{<<([]<>){<><>}>[<(){}>]><{[<>[]](()<>)}[{()()}]>}<({{{}{}}}{[{}{}]<<>()>})<(<[]<",
		"(<({[{{[[(<[<(<>[])[[][]]>]{((<>[])({}<>))({[]<>})}>({<[()[]]]<{()[]}>})){(<(<<>[]>[<>{}])>{[({}<",
		"(<{<(({[(({[[<<>{}>({}<>)]{({}{}){(){}}}]{(<(){}>[[][]])<{()<>}>}}[[[{[]{}}[()]]](([[]()][(){}]))])",
		"([({<{[{<[{<({{}()}{{}()})>}]>}<<[[{{((){})<<>{}>}(({}[]){<>[]})}]]><[{{<({}<>)(()())>[{<>()}{{}()}]}{<{",
		"<[{[({[<<[[<<{{}{}}<{}{}>>([<>][(){}])><<<{}{}>>>]]{([{[[]()][[][]]}<{[]<>}>]<<[<>[]][[]()]><<()[",
		"[{({{{<<{[{([[()]({}[])]{[()()]})(((<>{})[[][]])<{{}{}}>)}]}>>}[((({{<([[]()][[]])}}{{<({}[])[{}",
		"{[({<{<<{[(([<<>[]>({}{})>))]{(<{{{}<>}<<>{}>}>)}}><(<(({([]())(<>[])}{{[]()}{[][]}})<{(()[])[<>{}]}>){({",
		"[<((<[<<(<<{<(<>)<[]{}>>[<[]{}><[]()>]}<((()[])((){}))[({}[]){[]{}}]>>([<[{}()][[]{}]>{{()[]}}][{[{}<>]<[]<",
		"[<[({({[([{<(([]<>)({}[]))<{{}[]}[<>[]]>><[<{}[]>({}<>)]<(())>>]]([[[{[]<>}({}[])](<[]{}>{<>()",
		"{[<<{(<[{(<[({[]{}}{[]{}}}]({<()[]><<>[]>}{<[]()>[{}{}]})>)}{[{(((<><>))(<[]<>>(()[])))[[<[]()>[{}()]](((",
		"({[<<([[{(<<<{<>()}<{}<>>>[[[]{}){()}]><[[(){}]<[]<>>]({(){}}[[]{}])>>{[{[()<>][{}()]}][[[{}{}]<<>()",
		"([(<((<{[[[[[{{}()}{()[]}]]<([[]<>])<[()<>][[][]]>>]][{([{()[]}(()[])]({<>{}}{[]()}))}({([[]",
		"(({{[<<{(([{<{()()}[[]()]>{<()<>>[[]{}]}}{[<()()>[[]{}]]{([]())({}[])}}][(<{<>[]}[()()]><<()>(()())>)",
		"[<(<{{{<{(<{(<<>[]><<>()>){<{}()><()<>>}}(<[<>{}]{<>{}}><[{}{}][<>()]>)>)}{{{(<(<><>)(<>())><{[]}([][",
		"<<<{([[<[[[(([[]{}])<({}())(<>())>)({[<>{}]{[]()}}<<(){}>{(){}}>)]{([(()()>({}[])]((<><>)<",
		"{{({<<<[[<(<{{{}<>}({}[])}{<{}<>><<>()>}>{({{}<>}({}()}){[[]<>]([]{})}}){{[[[][]](()())]}([<<>{}",
		"<<[[<{{({<{[{<<>()><()()>}<[()<>]<()()>>]{(([]<>)[{}[]]){{{}<>}{[]<>})}}>{{<[<{}<>>({}())]",
		"({[<[[{<<{{[{(<>{})}]{<[{}{}][[]()]>[[[]{}][()()]]}}[(<<{}[]>{<>[]}>{{{}<>}[(){}]}){{[<>{}]<{}{}>}{",
		"(([<{<[[([[(((()<>))({{}[]}<()<>>))[[({}[])(())]<{(){})>]][{((()){[]()}){{()<>}{<><>}}}[{{()}{()[]}}]]]){{(({",
		"{<[{(<[{[{{[<([]<>)<{}<>>>]}}]}]><{{<[[([{<>()}[[]{}]]({{}()}{<>{}}))]{([<[]{}>(()<>)]<[[]<>]([])",
		"{[<<[<(<[{<<[{{}()}{<>{}}](<{}<>>({}<>))>><{<({}[])<()<>>>[[<>[]][<><>]]}[<(<><>){{}<>}>[<[]()>{[]}]]>}]{[<{<",
		"(<({{[{<(<<<{[{}<>]({}[])}{[[]{}]<()[]>}>(<[[]()][()()]><[()[]][<>()]>)>>[(<<<{}()>{<><>}>>){<<{<><>}(",
		"(<{[[{{<(<(<({{}<>}{{}()}){[()<>]<{}[]>}>[(<{}{}><[]()>)<({}<>){()()}>]){[[([]{}){{}[]}]<[{}{}]<[][]>>][{<()",
		"[(<([({[[[<[[[[][]][{}{}]]([{}()]{<>{}})]{<{()()}<()<>>><{{}<>}<(){}>>}>]{{[([[][]]<()()>)",
		"(<(<<({{(<[{{[()<>][<>{}]}[<[]<>>]}<<{[]{}}>[[{}()]{()()}]>]{[([()()]<<>>)({()<>}[<>[]])][{{[][]}(<>())}",
		"[[((({(({{{[{{<>[]}[<>{}]}[[{}<>]([]())]][{(<><>)}]}][(<{{[][]}(<><>)}<[[]()]>>([[{}()]<[]<>>][<[][]>[()()]])",
		"<{({(([((<[<{(()[])<[][]>}{<<><>><[]<>>}><<([]<>)([]<>)>>]{[<([]{}]{<>{}}>]}>[{[(([]{})[{}{}])[<<>{}",
		"{(([<{(({<({{{{}{}}}[[[]]{[]()}]}{<[[]]<<>()>>(<()>[()<>])})>{{[(([]<>){{}()}){[{}]{{}<>}}]}([[([]{})]<[[](",
		"<{[({<{[[[{{{[{}{}]}}([<[][]>[<>[]}][<<>()>{()}])}((<({}[])[<><>]>(((){})[{}()]))[[<{}[]>[(",
		"{<[([(<((<<<<[{}[]]({}{})>((()())[(){}])><([[]{}][()<>])<<<>[]>([]<>)>>><[[((){}){[][]}}{[[]]",
		"[{(<{<{<<({({(()<>){{}[]}}([[]()]{()<>}))(([[]<>](())))}<{({()[]}{[][]})}])>>}>(({[{[{{([]())(<><>)",
		"<[[<([<<({[[(<[]()><<><>>)({<>})]{{<<><>>[()[]]}([()<>](()[]))}]((<{<>{}}<(){}>>))}<<[(<(){}><<>{}>)[[()[",
		"<[{<<<{[{<{(<[<>{}][{}()]>((<>[])[<><>]))<[({}<>)<<>[]>]<(<><>)(<>{})>>}({<<[][]>([]<>)>([<>{}><<>[]>",
		"[<<(<{((<([<({{}{}](<>[]))[[()<>]([][])]>((([])<(){}>))]){<{{<()[]>(()<>)}(<[]<>>([][]))}><[([(",
		"{[[[{{<[[{<[<<<>[]>{<><>}>{<<>{}><[]<>>}]((<{}()>(<>[])))><<{<<>{}>{(){}}}([{}{}]<<>()>)><[{<>{}}]>>}]",
		"((((([[(([{<({()[]}<()<>>)><<<<>()>({}{})>{{(){}}(<>[])}>}<<<{<><>}>><(((){})<[]()>){{()()}({}<>)}>",
		"<<[<[{<[<[<{[[()[]]({}[])]{<{}<>>[{}{}]}}[([[][]])[([]())[[]<>]]]>[(<<{}()>(()<>)>{{[]{}}<{}{}>})({([",
		"[{<(<{(<[<[[[{()<>}[[][]]][(()<>){{}[]}]](<{<>{}}[()[]]>(<[]>[{}()]))]([([<>{}]<()[]>){{<>{}}{()",
		"<[<({<({[{({{(<>{}){[]()}}<{[]<>}[(){}]>}({[<>[]]{[]<>}}))<{([{}[]])(({}[]){(){}})}[{[[]<>]}[({}())[(){}]]]",
		"[[[{<[((<[<{([[]<>]{<><>})<<{}[]>[[][]]>}{[[<>{}]]<<{}>([]<>)>}>{<{{<>[]}}[(<><>)[[]<>]])<{[(){}]<()()>}>}]{{",
		"([(<({[[{<([{{<>{}}([]<>)}[[{}[]]{<>{}}]](<[{}<>]{[]()}><((){})[[][])>)){({{{}{}}{{}{}}}[<{}{}>({}<",
		"<[{(((({[[{{({{}[]]<<>>)<{[]()}<()<>>>}((<(){}>([]{})))}{({[()()]([]<>)}<([]<>)[()()]>)<[{{}[",
		"<[({[<<<<{[({{(){}}[{}<>}}(<<>[]>[[]()])){(<()[]><<>()>)}]}><({<[{()()}]({[]{}}<{}>)>[({[]",
		"<<([[<{{([{<{([]{})[(){}]}{([]{})<{}[]>}><<(<>{})[<>()]>({{}<>})>}<[(<<>()><{}()>)<{()<>}<<>>",
		"<[[[({[({<[{((<>[]]<{}<>>){<<>()>{[]()}}}<{<{}{}>[<>{}]}{{{}[]}{{}()}}>]{<{{()()}<[]<>>}((",
		"[[({(([[(([<[([]<>){{}()}]>(<{{}{}}(()<>)>[{<><>}{{}<>}])]{{<[[]<>]>[[{}()]]}([<{}()>][([]{})<{}",
		"<(([<{([[<[((([]<>)<<><>>){[[]()]({}{})})({{(){}}((){}]}({{}()}))]([<{<>()}<[]{}>><{[]{}}<(){}>>])>",
		"<<[[{<<{[(<<{{[][]}[[]<>]}<<[][]><<>[]>]>([{<>}[<>[]]]{{[]()}<{}()>})>{({({}<>){[]<>}}[{<>{}}[<>[]]",
		"((<<<<[[(({[[<()<>>[(){}]]{(<>()){<>{}}}]}{(([<>[]><<>()>)<<(){}><()()>>)[(<{}{}>((){}))[{[][]}(()<>)]]})){",
		"[({(<{{<<<{<[{[]()}<<><>>][[()[]][<>[]]]>}>([([<[]<>>]{<<>()><()<>>})({<{}[]><()()>}[[[]()]])]<{<(()",
		"([(<({[([{{{((()())[<>[]]){{[]<>}[{}{}]}}([{()<>}<{}()>])}}<{[({(){}}(<><>)}<(<>())>][<[(){}]{(){}}",
		"{[[{((((([[{{[[]{}]({}{})}{[[]<>](()[])}}[{[()()]}((<>{})<(){}>)]]{(<{<>()}<(){}>>){(<{}()>{{}<>})({<>(",
		"([[{<[<[[<<{<{()<>}[<>()]>[[[]{}]{{}{}}]}<[{{}[]}]>>{<{{<>{}}([]())}>[(<<>()>)(<{}()>[<>()])]}>{({<[()<>]",
		"{(([{[{{(({(<[()[]][<>[]]>)<[(()<>){{}{}}]>}[(<<{}{}><{}<>>>){<[()()]<()()>>({[][]}<<><>>)}",
		"<[{<(<<[{[({[({}<>)<<>[]>](<<>()>(<><>))})[([{<><>}<()()>])]]}]><<[<({{{<>}}<([]][<>{}]>}[<(())<<><>>>])",
		"(<(({<[[{{[<[{<>[]}[()<>])((()()){{}[]})>]}}]]((<<<((({}{})<{}>){{{}<>}{<>}})[({{}()}{<>()}){{{}()}}]",
		"{[<[{(([<[{{[([]{})<[][]>]{([][])[()<>]}}<<(()<>)>>}[[(({}{}){{}<>})]]]{{[(<{}{}>[{}()])({()()}",
		"[{{((<<[[((<<{[]<>}<{}>>(<{}<>>{[]()})>({<()>([][])><[{}{}]{{}()}>))(({([]<>)([]())}<{[][]}<()>>)",
		"({([{({[({[<({{}{}}([]()))(((){}){{}{}}>>({(()()){{}()}})]<[[{{}{}}<{}[]>]{<[][]><<>[]>}]({<<>[]>",
		"{[({[{{[<[{([<{}{}>[(){}]](<{}<>>{{}()}))([([]()){()[]}]([{}<>]{<>{}}))}]>[{{[[{[][]}]]<((()[])([]{}))<[()",
		"[[(({{{(([[[{[()()]([]())}((<>){{}{}])]]({<[<>]({}{})>([<>{}]{{}{}})})]))}[(<[[(([[][]]{<>{}}){{(){}}})",
		"<(<({<{<[<(<{<()><<><>>}<<{}[]>([]<>)>>)<{<{[]{}}[{}()]>}([(()[]){[]<>}]{{<>()}{{}<>}})>}{[[{(()[]){(",
		"[(<(<([{<[{<(<()>([]<>))[{()()}<{}<>>]><{([]())<<>()>}(<{}{}>[()[]])>}]<{<({[][]}<[]()>){{{}[",
		"[<({<{(<<[[[<{()[]}{{}[]}>{[<><>]}]{[<()<>>({}<>)]}]][[{<[<>[]>([]())><{()()}{[]{}}>}([({}{})[[](",
		"<<[[{{<<([[{[(()())[()]][[{}[]]{{}}]}[[({}<>){[]{}}]{{<><>}[()[]]}]]{<[(<>())]>{({{}()}(()<>))}}])",
		"{<[{({[({{<({[{}()]([][])}([{}<>]{[]()}))[<(<><>)([]{})><{()()}(<>[])>]>{{(<[]<>>[{}])}[<{(){}}(<>",
		"<{{(({([[<[{<(<>)([]{})>{<<>{}>{(){}}}}]{[<([]{})<{}{}>>[[<>{}]{()<>}]][{{<><>}([]())}({()<>}<()>)]}>]",
		"{<{[{[({[(({<[{}()]({}[]}>[<()>]}{(<<>>)}))]})]{<{<{[[{(<>())[()[]]}[[{}{}][[]]]][([[]()]<{}<>>)[<[]{}>([]<>",
		"((<{{<[[<{{([([]{})<<>()>]<(<>{})(()<>)))([{()()}(<><>)]<(<>)(<>{})>)}([<{(){}}<<>{}>>]<[([]<",
		"[<<<{(((<<{([<()[]>({})]({<>()}(()())))({([])<<>>}((()[])[()[]]))}{<[<{}()>]<[{}{}]({}<>)>>{<<",
		"[[[<<<<[{(<<[{{}[]}]<(<><>)(<>[])>><<<<><>][[][]]>(([]{}))>>([<[[]<>]<[]<>>><{()<>}([][])>]))}<(",
		"<<{[[[{{{{[[(([]{}))([<>[]]<()>)][{(<>()){[]()}})]{(({<><>}{<>[]})[<<>[]>[()()]])}}}<{(<((<><>)(<>{}))<{",
	}

	totalSyntaxErrorScore := 0
	encounteredOpeningCharacters := make([]byte, len(inputRows[0]))
	completionScores := make([]int, len(inputRows))
	completionScoresCount := 0
	for _, inputRow := range inputRows {
		encounterIndex := -1
		corrupted := false
		for i := 0; i < len(inputRow); i++ {
			char := inputRow[i]
			if char == '(' || char == '[' || char == '{' || char == '<' {
				encounterIndex++
				encounteredOpeningCharacters[encounterIndex] = char
				continue
			}

			encounterIndex--
			if char == ')' && encounteredOpeningCharacters[encounterIndex + 1] != '('{
				totalSyntaxErrorScore += 3
				corrupted = true
				break
			}
			if char == ']' && encounteredOpeningCharacters[encounterIndex + 1] != '[' {
				totalSyntaxErrorScore += 57
				corrupted = true
				break
			}
			if char == '}' && encounteredOpeningCharacters[encounterIndex + 1] != '{' {
				totalSyntaxErrorScore += 1197
				corrupted = true
				break
			}
			if char == '>' && encounteredOpeningCharacters[encounterIndex + 1] != '<' {
				totalSyntaxErrorScore += 25137
				corrupted = true
				break
			}
		}

		if corrupted {
			continue
		}

		completionScore := 0
		for i := encounterIndex; i >= 0; i-- {
			completionScore *= 5
			switch encounteredOpeningCharacters[i] {
			case '(':
				completionScore += 1
			case '[':
				completionScore += 2
			case '{':
				completionScore += 3
			case '<':
				completionScore += 4
			}
		}
		completionScores[completionScoresCount] = completionScore
		completionScoresCount++
	}

	for i := 0; i < completionScoresCount; i++ {
		for j := i + 1; j < completionScoresCount; j++ {
			if completionScores[j] < completionScores[i] {
				temp := completionScores[j]
				completionScores[j] = completionScores[i]
				completionScores[i] = temp
			}
		}
	}

	fmt.Println("Total syntax error score :", totalSyntaxErrorScore)
	fmt.Println(len(completionScores), "completion scores, middle is", completionScores[completionScoresCount/2])
}