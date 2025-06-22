export function debounce(func: Function, wait: number) {
  var timeout: NodeJS.Timeout | undefined;

  return function executedFunction(this: any) {
    var context: any = this;
    var args = arguments;

    var later = function () {
      timeout = undefined;
      func.apply(context, args);
    };

    clearTimeout(timeout);

    timeout = setTimeout(later, wait);
  };
};

export class BetterSet<v> extends Set<v> {
  public toArray(): v[] {
    let arr: v[] = new Array<v>();

    for (var item of Array.from(this.values())) {
      arr.push(item);
    }

    return arr;
  };

  public join(str: string): string {
    return this.toArray().join(str);
  }
}