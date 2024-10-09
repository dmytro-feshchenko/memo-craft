export interface TransoformTextEngineOptions {

}

export interface TransformTextToHtmlEngine {
    transformToHtml(text: string, options?: TransoformTextEngineOptions): string;
}