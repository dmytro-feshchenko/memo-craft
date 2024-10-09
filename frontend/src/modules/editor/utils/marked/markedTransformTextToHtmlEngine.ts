import { Marked } from 'marked';
import { markedHighlight } from 'marked-highlight';
import Prism from 'prismjs';
import { TransformTextToHtmlEngine, TransoformTextEngineOptions } from '../types';
import { DEFAULT_TEXT_TO_HTML_TRANSFORM_OPTIONS } from './options';

export class MarkedTransformTextToHtmlEngine implements TransformTextToHtmlEngine {
    private engine: Marked;

    constructor() {
        const self = this;
        this.engine = new Marked(
            markedHighlight({
                highlight(code, lang) {
                    return self.parseSyntaxLanguageGrammar(code, lang);
                },
            }),
        );
    }

    /**
     * Parses the given code block and uses Prismjs to highlight it
     * @param code Code to parse
     * @param lang Language from https://prismjs.com/#supported-languages
     * @returns 
     */
    parseSyntaxLanguageGrammar(code: string, lang: string): string {
        if (!lang) return code;

        const grammar = Prism.languages[lang];
        if (!grammar) {
            console.error(`Prism can't process grammar for the language: "${lang}". Check supported languages on`
                + `https://prismjs.com/#supported-languages`
            );
            return code;
        }
        return Prism.highlight(code, grammar, lang);
    }

    transformToHtml(text: string, options: TransoformTextEngineOptions = {}): string {
        options = Object.assign({}, DEFAULT_TEXT_TO_HTML_TRANSFORM_OPTIONS, options);

        let resultHtml = '';

        

        return resultHtml;
    }
}