import { Marked } from 'marked';
import { markedHighlight } from 'marked-highlight';

const marked = new Marked(
    markedHighlight({
        highlight(code, lang) {
            // Language may be undefined (GH#591)
            // if (!lang)
                return code;

            // if (DIAGRAM_TYPE.includes(lang))
            //     return code;

            // const grammar = Prism.languages[lang];
            // if (!grammar) {
            //     console.warn(`Unable to find grammar for "${lang}".`);
            //     return code;
            // }
            // return Prism.highlight(code, grammar, lang);
        },
    }),
);

export class DataTransformer {
    private 
    constructor() {

    }
}