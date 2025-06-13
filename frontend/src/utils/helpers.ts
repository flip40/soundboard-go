export function debounce(func: Function, wait: number) {
  // console.log("debounce");
  var timeout: NodeJS.Timeout | undefined;

  return function executedFunction(this: any) {
    // console.log("executedFunction");
    var context: any = this;
    var args = arguments;

    var later = function () {
      // console.log("later");
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