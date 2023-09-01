package javanalyser

import (
	"unicode"

	"fr.cybercicco/springgo/spring-cli/entities/enums"
)


type SyntaxToken struct {
    Value string
    position int16
    kind string
}

var position int16
var line []rune

func next() string{
    lastPosition := position
    position++
    return string(line[lastPosition])

}

func lex(line []rune) SyntaxToken{
    startPos := position
    if unicode.IsLetter(line[position]) {
        position++
        for unicode.IsLetter(line[position]) || unicode.IsDigit(line[position]){
            position++
        }
        token := SyntaxToken{ Value : string(line[startPos:position]) ,position : startPos, kind : enums.WORD_KIND, }
        return token
    }
    if unicode.IsDigit(line[position]) {
        position++
        for unicode.IsDigit(line[position]) || line[position] == '.' {
            position++
        }
        token := SyntaxToken{ Value : string(line[startPos:position]) ,position : startPos, kind : enums.NUMBER_KIND, }
        return token
    }
    if line[position] == '"' && position != 0 && line[position-1] != '\\' {
        position++
        for line[position] != '"' {
            position++
        }
        token := SyntaxToken{ Value : string(line[startPos + 1 :position]) ,position : startPos, kind : enums.STRING_KIND, }
        position++
        return token
    }

    if line[position] == '/'{
        if line[position+1] == '/' {
            for position+1 != int16(len(line)) && line[position] != '\n'{
                position++
            }
            return lex(line)
        }
        if line[position+1] == '*'{
            for position+1 != int16(len(line)) && (line[position] != '*' || line[position+1] != '/'){
                position++
            }
            position += 2
            return lex(line)
        }
    }

    switch line[position] {
    case ' ':
        position++
        return lex(line)
    case '\n':
        position++
        return lex(line)
    case '@':
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.ANNOTATION_DELIMITER_KIND, }
    case '(':
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.OPEN_PARENTHESIS_KIND, }
    case ')':
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.CLOSE_PARENTHESIS_KIND, }
    case '{':
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.OPEN_BRACKET_KIND, }
    case '}':
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.CLOSE_BRACKET_KIND, }
    case '<':
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.OPEN_TYPE_KIND, }
    case '>':
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.CLOSE_TYPE_KIND, }
    case ';':
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.END_OF_LINE_KIND, }
    case '*':
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.STAR_KIND, }
    case '.':
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.DOT_KIND, }
    case ',':
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.COMMA_KIND, }
    case '=':
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.EQUAL_KIND, }
    default:
        return SyntaxToken{ Value  : next(), position : startPos, kind : enums.BAD_TOKEN, }
    }


}

func lexLine(newLine []rune) []SyntaxToken{
    line = newLine
    syntaxTokens := []SyntaxToken{lex(line)}
    for syntaxTokens[len(syntaxTokens)-1].kind != enums.END_OF_LINE_KIND && syntaxTokens[len(syntaxTokens)-1].kind != enums.CLOSE_BRACKET_KIND && syntaxTokens[len(syntaxTokens)-1].kind != enums.OPEN_BRACKET_KIND && syntaxTokens[len(syntaxTokens)-1].kind != enums.COMMENTARY_KIND && int(position) < len(newLine)-1 {
        syntaxTokens = append(syntaxTokens, lex(line))
    }
    return syntaxTokens
}

func LexFile(lines *string) [][]SyntaxToken{
    position = 0
    tokens := [][]SyntaxToken{}
    for position < int16(len([]rune(*lines))-1){
        tokens = append(tokens, lexLine([]rune(*lines)))
    }
    return tokens
}
