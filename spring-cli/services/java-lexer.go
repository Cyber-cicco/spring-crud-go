package services

import (
	"unicode"
)

var WORD_KIND = "WORD_KIND"
var SPACE_KIND = "SPACE_KIND"
var END_OF_LINE_KIND = "END_OF_FILE_KIND"
var COMMMENT_DELIMITER_KIND = "COMMENT_DELIMITER_KIND"
var STRING_DELIMITER_KIND = "STRING_DELIMITER_KIND"
var ANNOTATION_DELIMITER_KIND = "ANNOTATION_DELIMITER_KIND"
var OPEN_PARENTHESIS_KIND = "OPEN_PARENTHESIS_KIND"
var CLOSE_PARENTHESIS_KIND = "CLOSE_PARENTHESIS_KIND"
var OPEN_BRACKET_KIND = "OPEN_BRACKET_KIND"
var CLOSE_BRACKET_KIND = "CLOSE_BRACKET_KIND"
var OPEN_TYPE_KIND = "OPEN_TYPE_KIND"
var CLOSE_TYPE_KIND = "CLOSE_TYPE_KIND"
var END_OF_LINE_TOKEN = "END_OF_LINE_TOKEN"
var BAD_TOKEN = "BAD_TOKEN"
var DOT_TOKEN = "DOT_TOKEN"

type SyntaxToken struct {
    value []rune
    position int16
    kind string
}

type SyntaxFile struct {
    tokens [][]SyntaxToken
}

var position int16
var line []rune

func next() rune{
    lastPosition := position
    position++
    return line[lastPosition]

}

func lex(line []rune) SyntaxToken{
    startPos := position
    if unicode.IsLetter(line[position]) {
        position++
        for unicode.IsLetter(line[position]) || unicode.IsDigit(line[position]){
            position++
        }
        token := SyntaxToken{ value : line[startPos:position] ,position : startPos, kind : WORD_KIND, }
        return token
    }

    if line[position] == '/'{
        if line[position+1] == '/'{
            token := SyntaxToken{ value : line[startPos:position] ,position : startPos, kind : WORD_KIND, }
            return token
        }
    }

    switch line[position] {
        case ' ':
            return SyntaxToken{ value  : []rune{next()}, position : startPos, kind : SPACE_KIND, }
        case '@':
            return SyntaxToken{ value  : []rune{next()}, position : startPos, kind : ANNOTATION_DELIMITER_KIND, }
        case '"':
            return SyntaxToken{ value  : []rune{next()}, position : startPos, kind : STRING_DELIMITER_KIND, }
        case '(':
            return SyntaxToken{ value  : []rune{next()}, position : startPos, kind : OPEN_PARENTHESIS_KIND, }
        case ')':
            return SyntaxToken{ value  : []rune{next()}, position : startPos, kind : CLOSE_PARENTHESIS_KIND, }
        case '{':
            return SyntaxToken{ value  : []rune{next()}, position : startPos, kind : OPEN_BRACKET_KIND, }
        case '}':
            return SyntaxToken{ value  : []rune{next()}, position : startPos, kind : CLOSE_BRACKET_KIND, }
        case '<':
            return SyntaxToken{ value  : []rune{next()}, position : startPos, kind : OPEN_TYPE_KIND, }
        case '>':
            return SyntaxToken{ value  : []rune{next()}, position : startPos, kind : CLOSE_TYPE_KIND, }
        case ';':
            return SyntaxToken{ value  : []rune{next()}, position : startPos, kind : END_OF_LINE_KIND, }
        case '.':
            return SyntaxToken{ value  : []rune{next()}, position : startPos, kind : DOT_TOKEN, }
        default:
            return SyntaxToken{ value  : []rune{next()}, position : startPos, kind : BAD_TOKEN, }
        }
        

}

func lexLine(newLine []rune) []SyntaxToken{
    line = newLine
    syntaxTokens := []SyntaxToken{lex(line)}
    for syntaxTokens[len(syntaxTokens)-1].kind != END_OF_LINE_KIND && syntaxTokens[len(syntaxTokens)-1].kind != CLOSE_BRACKET_KIND && syntaxTokens[len(syntaxTokens)-1].kind != OPEN_BRACKET_KIND {
        syntaxTokens = append(syntaxTokens, lex(line))
    }
    return syntaxTokens
}

func LexFile(lines *string) SyntaxFile{
    syntaxFile := SyntaxFile{}
    for position < int16(len([]rune(*lines))-1){
        syntaxFile.tokens = append(syntaxFile.tokens, lexLine([]rune(*lines)))
    }
    return syntaxFile
}
