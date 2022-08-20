local arr = std.range(2, 1000);
{
  fields: {
      //fieldname : fieldvalue
      ['f' + x]: x
      for x in arr
  },
}