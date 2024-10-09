import { IMetaData } from "../../../models/metadata";
import generateUUID from "../../../utils/generateUUID";

export class FlashcardModel implements IMetaData {
    id: string = generateUUID();
    createdAt: number = Date.now();
    updatedAt: number = Date.now();
}